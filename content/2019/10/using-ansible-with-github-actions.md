---
title: Using Ansible with GitHub Actions
date: "2019-10-21"
twitter_card:
  description: "Running Ansible playbooks using GitHub Actions"
  image: "https://misfra.me/img/2019/10/github-actions-ansible.png"
---

My most recent work with GitHub Actions involves migrating [Transverse](https://transverseapp.com) from being
deployed manually using Ansible run on my laptop to a CI/CD approach using Ansible on GitHub Actions.
Now I can push changes and deploy from anywhere without requiring access to my personal laptop or the private
keys to connect to my server.

<!--more-->

_Note:_ Transverse is open source so you can look at the actual GitHub Actions files in the [repo](https://github.com/Preetam/transverse).
I've copied the relevant portions here for convenience.

### How it's built and run

Transverse runs as two containers, but both are run from the same Docker image with different parameters. The image
is built using `docker build` and the image is pushed to Docker Hub. This is also done using GitHub Actions and
I've described that workflow in a [previous post](/2019/09/deploying-projects-to-cloud-run-using-github-actions/).

The Docker containers run on a DigitalOcean droplet as [runit](http://smarden.org/runit/) services. A "deploy" in this case just means
SSHing into the server, updating the runit service scripts with the new Docker image tags, and restarting the services. This is
done with an Ansible playbook.

### Implementation

I couldn't find a GitHub Action to run Ansible playbooks so I ended up creating my own. You can define custom Actions as Docker images.
I created one specifically for my project and it's located in
[.github/actions/ansible](https://github.com/Preetam/transverse/tree/428d9ee4947cc359fe91730a908bd08771e48b0b/.github/actions/ansible).

This Docker image is built and run as part of my workflow on every merge to master. The Dockerfile is a minimal Alpine image with
Ansible. I had to experiment with the dependencies and pip in order to get my playbook running. Here's what I ended up with.

**Dockerfile:**

```dockerfile
FROM alpine 

ENV ANSIBLE_HOST_KEY_CHECKING=False

RUN apk add ansible gcc python3-dev libc-dev libffi-dev openssl-dev
RUN pip3 install --upgrade paramiko
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

The entrypoint script is where the interesting stuff happens. My Ansible image needs to be able to connect to my
DigitalOcean droplet over SSH, and I somehow need to _securely_ provide an SSH private key to the container. What I
ended up doing is using Ansible Vault to encrypt a private key I created specifically for GitHub Actions, and checking
it into the repo.

In the entrypoint script, I run `ansible-vault` with the vault password passed in as an environment variable to
decrypt the SSH key and prep it to be used with `ansible-playbook`. The vault password is set in GitHub as
a secret environment variable.

**entrypoint.sh:**

```
#!/bin/sh

echo "$VAULT_PASS" > ~/.vault_pass.txt
mkdir ~/.ssh
ansible-vault --vault-password-file ~/.vault_pass.txt view ansible/ssh_key.txt.secret > ~/.ssh/id_rsa
chmod 0600 ~/.ssh/id_rsa

ansible-playbook -e "build_sha=$GITHUB_SHA" --vault-password-file ~/.vault_pass.txt -i ansible/hosts ansible/deploy.yml
```

And that's it! The last thing to do is include the Ansible action as part of the workflow.

**.github/workflows/master_build_and_deploy.yml:**

```yaml
  deploy:
    runs-on: ubuntu-latest
    needs: build
    steps:
    - uses: actions/checkout@v1
    - uses: ./.github/actions/ansible
      env:
        VAULT_PASS: ${{ secrets.VAULT_PASS }}
```

![GitHub Actions screenshot](/img/2019/10/github-actions-ansible.png)

Finally, I can fix and deploy typos without having to wait until I have access to my personal laptop. 🙂
