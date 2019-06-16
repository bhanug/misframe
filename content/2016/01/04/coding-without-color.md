---
title: Coding without Color
date: "2016-01-04T19:47:55.996Z"
---

[Jasdev requested](https://twitter.com/jasdev/status/672795506242953216) this post a month ago and
I've just gotten around to it. As the title suggests, my development environment is almost entirely
monochromatic. I use the default Terminal theme on OS X and made my own theme for Sublime Text 3.

My Sublime theme Infimum is available on [GitHub](https://github.com/Preetam/Infimum) and I wrote
a [post](/syntax-highlighting/) about it before. Creating a theme is very simple. I used an online
editor called TmTheme Editor. It's not great but it has allowed me to make and view changes quickly.
You can load a tmTheme file directly from a URL, so you can [play around](https://tmtheme-editor.herokuapp.com/#!/editor/url/https://raw.githubusercontent.com/Preetam/Infimum/master/Infimum.tmTheme) with Infimum from the
GitHub repository. Many of the scopes are assigned the default style: black on white without any
text decoration.

![](/img/2016/01/04/coding-without-color/tmtheme-editor.png)

After using this theme for over a year, I decided that many syntax highlighting themes use colors
excessively. They remind me of Winamp skins and Windows XP themes from way back. I don't find them
cool anymore.

My syntax highlighting settings offer minimal guidance. They help me find rough structure in my code
and the occasional typos. Most importantly, they're not visually distracting, unlike some from
previous themes I used.

As the Sublime default, Monokai ended up being the theme I used when I first started using the
editor. I tried it out again and looked for things that did not appeal to me. What stands out for me
the most in the following image is how the header file has member function names highlighted in
green but the source file does not. I don't like those inconsistencies.

![Monokai](/img/2016/01/04/coding-without-color/monokai.png)

Solarized is (was?) another popular theme. I heard about it during high school (~8 years ago), and
thought it was technical and fancy. I don't think it provides enough contrast for me now.

![Solarized](/img/2016/01/04/coding-without-color/solarized.png)

I was surprised to find that it *also* has some weird inconsistencies. Take a look at the braces in
the following image:

![Solarized CSS](/img/2016/01/04/coding-without-color/solarized-css.png)

Why are two colors used? I just don't get it.

Infimum isn't perfect yet. I give it tweaks every now and then, but those have been rare. I don't
really notice it anymore, which was the entire point!

---

## Hex table for reference

| Scope | Hex |
|-----------------------|------------------------------|
| Comment | #999999 |
| Number | #000000 |
| Built-in constant | #333333 |
| User-defined constant | #000000 |
| Keyword | #000000 |
| Storage | #000000 |
| Storage type | #000000 |
| Class name | #000000 |
| Inherited class | #000000 |
| Function name | #000000 |
| Function argument | #000000 |
| Tag name | #555555 |
| Tag attribute | #000000 |
| Library function | #000000 |
| Library constant | #000000 |
| Library class/type | #000000 |
| Invalid | #FFFFFF (background #000000) |
| Invalid deprecated | #474747 (background #E0E0E0) |
