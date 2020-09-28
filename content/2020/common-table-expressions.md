---
title: Common Table Expressions
date: "2020-09-28"
---

I use common table expressions (CTEs) in SQL queries a lot these days. CTEs allow
you to temporarily use the results of one query as a table in other queries.
You use them like this:

```sql
WITH cte AS (SELECT ...)
SELECT * FROM cte;
```

When would you want to use CTEs? One case is when you want to use a subquery, that you're already using as a
column, as a WHERE condition. Here's an example that uses a schema inspired by GitHub issues.

<!--more-->

```sql
CREATE TABLE issues (
	id BIGINT NOT NULL PRIMARY KEY,
	title TEXT NOT NULL,
	content TEXT NOT NULL
);

CREATE TABLE labels (
	id BIGINT,
	text TEXT
);

CREATE TABLE issue_labels (
	issue_id BIGINT,
	label_id BIGINT,
	added_at TIMESTAMP
);
```

Let's say you want to get all of the issues that have label ID 5 as the latest label.
We can start with a query to select each issue ID and its latest label in a subquery:

```sql
SELECT
	id,
	(SELECT label_id
		FROM issue_labels
		WHERE issue_id = issues.id
		ORDER BY added_at
		DESC LIMIT 1
	) AS latest_label
FROM issues;
```

If you try to add a WHERE clause referencing `latest_label`, you'd see this:

```sql
SELECT
	id,
	(SELECT label_id
		FROM issue_labels
		WHERE issue_id = issues.id
		ORDER BY added_at
		DESC LIMIT 1
	) AS latest_label
FROM issues
WHERE latest_label = 5;
```

```
ERROR:  column "latest_label" does not exist
LINE 11: WHERE latest_label = 5;
```

You can get around this by using the subquery again in the WHERE clause.

```sql
SELECT
	id,
	(SELECT label_id
		FROM issue_labels
		WHERE issue_id = issues.id
		ORDER BY added_at
		DESC LIMIT 1
	) AS latest_label
FROM issues
WHERE (SELECT label_id
		FROM issue_labels
		WHERE issue_id = issues.id
		ORDER BY added_at
		DESC LIMIT 1) = 5;
```

But clearly this is much harder to read. You can imagine how complicated it would get if
there were more subqueries to filter on. On the other hand, here's the same query written
with a CTE:

```sql
WITH latest_labels AS (
	SELECT
		id,
		(SELECT label_id
			FROM issue_labels
			WHERE issue_id = issues.id
			ORDER BY added_at
			DESC LIMIT 1
		) AS latest_label
	FROM issues
)
SELECT id, latest_label
FROM latest_labels
WHERE latest_label = 5;
```

It's a lot easier to understand. One thing to keep in mind is that these two queries have exactly the same
query plan. Even though it's not obvious in the CTE version, the subquery is evaluated twice per row just
like the first version.

In general CTEs won't make your queries faster by themselves. Their results are temporary
and you can't index on them so you should be careful with larger tables. However I think they
help you write faster queries more clearly.

The other great thing about CTEs is that you can nest them, e.g.

```sql
WITH baz AS (
	WITH bar AS (
		SELECT * FROM foo;
	)
	SELECT * FROM bar;
)
SELECT * FROM baz;
```

I find this really useful when writing analytics-style queries for Grafana dashboards.
It's easy to start with one layer or CTE at a time and logically build up to the
final result you want.
