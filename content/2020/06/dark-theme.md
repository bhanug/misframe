---
title: Dark theme
date: "2020-06-15"
---

I added a dark theme to this blog over the weekend. Here's how you can do
something similar with just a few CSS updates.

The first thing I did was move all color codes to variables. This is useful
in general so you can define colors once and reuse them instead of copying
and pasting color codes everywhere.

```css
:root {
  --main-bg-color: white;
  --main-fg-color: black;
  --date-fg-color: #c0c0c0;
  --border-color: #eee;
}
```

Those are the only 4 colors I use. Next, I updated styles to reference those
variables, like this:

```
.mf-header-nav-links a {
  color: var(--main-fg-color);
}
```

I also had to add rules for things that assumed certain defaults, like the
`body` text color and background:

```css
body {
  color: var(--main-fg-color);
  background-color: var(--main-bg-color);
}
```

Finally, I added a media selector for dark color scheme preferences.
All it does is update the variables I defined earlier.

```
@media (prefers-color-scheme: dark) {
  :root {
    --main-bg-color: black;
    --main-fg-color: white;
    --date-fg-color: #888;
    --border-color: #333;
  }
}
```

And that's it! It only took a few minutes to do. All of the websites I have
created recently use CSS variables and the `prefers-color-scheme` media
selector because they're so useful.

<!--more-->
