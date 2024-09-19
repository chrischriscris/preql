# preql

A better thought out prequel of SQL.

## Why preql

Let's take a look at an example of a SQL query:

```sql
SELECT s.email, d.name, s.grade
FROM student s
JOIN department d
    ON s.department_id = d.id
WHERE (d.area = 'Math' OR d.area LIKE '% Science')
    AND s.grade > 16;
```

When you start writing the query your IDE can't really provide any autocompletion
because the result will mention tables that will be declared after (and a language
that makes you reference things that haven't been declare doesn't sound like a
smart design, isn't it?). The same query in preql would look like this:

```preql
query {
    tables {
        student := s
        department := d
    }

    joins {
        s.department_id = d.id
    }

    conds {
        d.area = 'Math', d.area like '% Science'
        s.grade > 16
    }

    get {
        s.email
        d.name
        s.grade
    }
}
```

Sure it isn't more concise but... it formats a lot more nicely, the code does really
writes itself and your IDE will have everything you need at the time you need it
to give you really nice autocompletions.

------
Made with :heart: by [chrischriscris](github.com/chrischriscris).
