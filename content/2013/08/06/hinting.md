---
title: Hinting
date: "2013-08-06"
url: /hinting
---


<p>Can you tell the difference between these two images? Click on them to see a full-size version.</p>
<p><a href="/img/copied/hinted.png" target="_blank"><img src="https://media.tumblr.com/58bee817767f5dd8bc91ac8cc1b16495/tumblr_inline_mr52atpwFa1qz4rgp.png" /></a></p>
<p></p>
<p><a href="/img/copied/unhinted.png" target="_blank"><img src="https://media.tumblr.com/328c56f4f684c3febbb8653df3da8c18/tumblr_inline_mr52b22bad1qz4rgp.png" /></a></p>
<p></p>
<p>Hint: there is a difference, and it's font hinting!</p>
<p>The first image has hinted text, and it's&nbsp;<em>terrible!</em>&nbsp;It's morphed. The dimensions are funky. Why destroy a perfectly fine typeface (created by&nbsp;<em>artists</em>) by hinting (using algorithms and fancy math)?</p>
<p>I turn hinting off everywhere. I suggest you do it too. If you use Ubuntu, you can install Unity Tweak Tool:</p>
<p><img src="https://media.tumblr.com/8ed9cda1202fb1d43c60f62585ce351d/tumblr_inline_mr52o4CCeB1qz4rgp.png" /></p>
<p>Sublime Text will still hint. You'll have to use a <code>~/.fonts.config</code> file to take care of that.</p>

	<?xml version="1.0"?>
	<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
	<fontconfig>
	<match target="font">
	   <edit name="hinting" mode="assign">
	      <bool>true</bool>
	   </edit>
	</match>

	<match target="font">
	   <edit name="hintstyle" mode="assign">
	         <const>hintnone</const>
	    </edit>
	</match>

	</fontconfig>

Make text look good!

