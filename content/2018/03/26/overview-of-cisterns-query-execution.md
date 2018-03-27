---
title: "Overview of Cistern’s query execution"
date: "2018-03-26T23:30:00-04:00"
---

I was sent an email asking about how Cistern's query execution works at a high level.
The author of the email does not have a strong computer science background and wanted to know
"the general steps involved in processing a query (i.e. what operations are performed in what order
and how they are performed at a high level)."

The code starts [here](https://github.com/Cistern/cistern/blob/15c114db0598e66780781a93585be3933454eb91/cmd/cistern/query.go#L49).

Here's my response that I typed up really quickly since I was already slow to respond. :)

---

At the outer layer, we first need to look at all of the rows within the time range requested. The key-value pairs are ordered by time so it's easy to seek to the first timestamp and limit the number of records read. The loop begins [here](https://github.com/Cistern/cistern/blob/15c114db0598e66780781a93585be3933454eb91/cmd/cistern/query.go#L88-L89).

Filtering is the first processing step after reading each record. If a record doesn't pass all of the filters, it's ignored and we continue to the next record.

One special case is when you don't have group-bys or aggregates. That's automatically considered to be a raw data query so all records that pass the filter are returned (up to a limit). It's just like `SELECT * WHERE ... LIMIT ...`.

Next comes grouping and aggregation. The data structure for those is basically map[string]float64 where the key is the grouping column values and the value and the aggregation value (count, min, max, sum). Time series aggregation also happens at the same stage because it's the same aggregation with a time bucket as the additional grouping column. This part is a lot like `WITH ROLLUP`, if you're familiar with that. Read more about it here: https://dev.mysql.com/doc/refman/5.7/en/group-by-modifiers.html

Finally, sorting and limiting of the groups happens. This can't happen before the end because if you order by `SUM(column_a)`, you have to have observed all of the possible values for `column_a`.

---

The code is really messy but the high level concepts are there if you look closely. I have an
[issue](https://github.com/Cistern/cistern/issues/90) to refactor all of that query execution code.
Once I have better abstractions I can implement things like secondary indexes!
