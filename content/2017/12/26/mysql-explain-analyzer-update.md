---
title: "MySQL Explain Analyzer update (new design and permalinks!)"
date: "2017-12-26T23:35:00-05:00"
---

tl;dr: Try it out here: https://preetam.github.io/explain-analyzer/#!/explain/

---

I [introduced my explain analyzer](/2017/11/19/mysql-explain-analyzer/) for MySQL a little over a month ago.
Here's what I wrote at the end of that post:

> It could look a lot better. The real thing isn’t as nice as the design, and the design is still pretty bad. I’ll work on that soon. In terms of the implementation, the explain analyzer is a single page app without any state, so if you refresh the page, you’ll lose everything. I think I’m going to add a “share” feature so that you can get a permalink with all of the values saved. That’ll take some more work but I think it would be a neat opportunity to use AWS Lambda for a user-facing site!

Both the design and the sharing feature have been addressed with this month's update!
Here's what the new design looks like:

![](/img/2017/12/explain-analyzer-update.png)

Click [here](https://preetam.github.io/explain-analyzer/#!/explain/3b0bb4c994eb6e79dccafb58682fe90860e41938.json) to
see that explain yourself! That explain output is from [this](https://www.percona.com/blog/2016/02/29/explain-format-json-nested-loop-makes-join-hierarchy-transparent/)
Percona post.

[Here](https://preetam.github.io/explain-analyzer/#!/explain/d14a5fe84e69f4b9fc4db9bd0ecbf9828d672f43.json) is another
explain from [another](https://www.percona.com/blog/2016/02/29/explain-format-json-nested-loop-makes-join-hierarchy-transparent/)
Percona post.

You'll notice that the comments for that are really repetitive. That's going to be addressed in the next
update.

> * Table "salaries": Matching rows are being accessed. MySQL is using the PRIMARY KEY. MySQL is using the 'PRIMARY' index.
> * Table "dept_manager": Matching rows are being accessed. MySQL is using the PRIMARY KEY. MySQL is using the 'PRIMARY' index.
> * Table "employees": At most one row is accessed from this table using an index. MySQL is using the PRIMARY KEY. MySQL is using the 'PRIMARY' index.

If you encounter any problems let me know on [GitHub](https://github.com/Preetam/explain-analyzer)
(with a permalink because now you can!). I'm always looking for more interesting explain plans.
