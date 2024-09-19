# Grammar of preql

- Example SQL query:

```sql
SELECT s.email, d.name, s.grade
FROM student s
JOIN department d
    ON s.department_id = d.id
WHERE (d.area = 'Math' OR d.area LIKE '% Science')
    AND s.grade > 16;
```

```preql
query {
    table {
        student := s
        department := d
    }

    join {
        s.department_id = d.id
    }

    cond {
        d.area = 'Math', d.area LIKE '% Science'
        s.grade > 16
    }

    get {
        s.email
        d.name
        s.grade
    }
}
```
