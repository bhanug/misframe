---
title: Heading hover hash links in JavaScript
date: "2016-12-30T13:00:00-05:00"
---

If the title doesn't make sense, maybe a GIF will.

<img src='/img/2016/12/heading-links.gif' width=369>

I wanted to add these to the [Epsilon docs page](https://epsilon.infinitynorm.com/docs/).
Here's how I did it using JavaScript and a little bit of CSS.

This is on the bottom of every page.

```js
// Find all headings under .content
var elems = document.querySelectorAll(".content h1, .content h2, .content h3, .content h4, .content h5, .content h6");

for (var i = 0; i < elems.length; i++) {
  var el = elems[i];
  var id = el.id;

  // Create the link
  var link = document.createElement("a");
  link.href = "#"+id;
  link.textContent = "#";

  // Add it to the heading element
  el.appendChild(link);

  // Set a couple of class names
  link.className = 'content-heading-link';
  el.className = 'content-heading';
}
```

Here's the CSS associated with those classes:

```css
.content-heading .content-heading-link {
    visibility: hidden;
    margin: 0 0.3rem;
}

.content-heading:hover .content-heading-link {
  visibility: visible;
}
```

It would be cool to be able to do this without JavaScript, but I'm not sure
how to do it using Hugo.
