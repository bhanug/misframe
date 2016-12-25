---
title: Simple CSS spinner
date: "2016-12-25T14:40:00-05:00"
---

Here's a really simple spinner that uses CSS:

![Spinner](/img/2016/12/spinner.gif)

All you need is an HTML element...

```html
<div class='spinner'></div>
```

...and some CSS.

```css
.spinner {
  height: 30px;
  width: 30px;
  border: 8px solid #8798A3;
  border-radius: 50%;
  border-left-color: #A4B1BA;
  animation: spin 0.75s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
```

Check it out on JSFiddle: https://jsfiddle.net/zjxske8z/
