---
title: Getting confused.
date: "2014-01-13"
url: /getting-confused
---


After working with C for a while, I feel like I understand C relatively less than before. I'm going to put this in the context of my `vlmap` ordered map data structure. Here's a code snippet to start off:

	uint8_t* key = "foo";
	int keylength = 3;
	uint8_t* val = "bar";
	int vallength = 3;

	vlmap* m = vlmap_create();
	vlmap_insert(m, key, keylength, val, vallength);

I've modified the original code for simplicity. It's simply inserting a key-value pair.

Here's my confusion: where do `key` and `value` live? On the stack, right? They're just fixed-length character arrays. I don't think I can just use those arrays directly in my data structure since, being stack-allocated, they'll disappear when the function ends. If this is indeed the case, then I'll have to do a `memcpy` into a `malloc`'d region of memory, and eventually I have to `free` it.

My memory's rather poor. I barely remember how some of those `vlmap` functions work, but that's the result of being *in the zone* late at night :). I wasn't copying over the strings initially and everything seemed fine. I thought it was because my entire test was running in `main()`, so I moved the key and value declarations into a function and set them in the map. They're still stack-allocated in that function, right? If they are, I would run the test and it should fail because those values would've been overwritten. I don't think it did. Huh?!

Then again, there might be a simple explanation for all of this, or I may be remembering things wrong. In any case, I'm not feeling confident about my C memory management skills!

