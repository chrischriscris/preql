I really like the direction you're taking with preql! The idea of declaring context before using it is fundamentally more ergonomic for both humans and tooling. Your approach reminds me of how modern query builders like Prisma or Drizzle work, but with a more SQL-native feel.

Here's how I see the strengths and potential extensions:

## What's Working Well

1. **Progressive disclosure** - You're building up the query context in a logical order
2. **Clear separation of concerns** - Each block has a single responsibility
3. **IDE-friendly** - Autocompletion can work naturally as context accumulates
4. **Implicit AND/OR semantics** - Using commas for OR within a group and separate lines for AND is intuitive

## Extending to Other SQL Constructs

Here are some ideas for expanding preql:

### Aggregations and Grouping
```preql
query {
    tables {
        order := o
        customer := c
    }
    joins {
        o.customer_id = c.id
    }
    group {
        c.country
        c.state
    }
    having {
        count(*) > 10
        sum(o.amount) > 1000
    }
    get {
        c.country
        c.state
        total_orders := count(*)
        revenue := sum(o.amount)
    }
}
```

### Subqueries as Named Blocks
```preql
query high_value_customers {
    tables { customer := c }
    conds { c.lifetime_value > 10000 }
    get { c.id, c.email }
}

query {
    tables {
        order := o
        high_value := @high_value_customers
    }
    joins {
        o.customer_id = high_value.id
    }
    get {
        high_value.email
        o.total
    }
}
```

### Window Functions
```preql
query {
    tables { sales := s }
    windows {
        monthly := partition by s.month order by s.date
        running := order by s.date
    }
    get {
        s.date
        s.amount
        month_rank := row_number() over monthly
        running_total := sum(s.amount) over running
    }
}
```

### Set Operations
```preql
query active_users {
    tables { user := u }
    conds { u.last_login > '2024-01-01' }
    get { u.id, u.email }
}

query premium_users {
    tables { subscription := s }
    conds { s.tier = 'premium' }
    get { s.user_id as id, s.user_email as email }
}

combine {
    active_users union premium_users
}
```

### CTEs (Common Table Expressions)
```preql
with {
    monthly_totals {
        tables { sales := s }
        group { month(s.date) as month }
        get { 
            month
            total := sum(s.amount)
        }
    }
    
    averages {
        tables { monthly_totals := mt }
        get {
            overall_avg := avg(mt.total)
        }
    }
}

query {
    tables {
        monthly_totals := mt
        averages := a
    }
    conds {
        mt.total > a.overall_avg
    }
    get {
        mt.month
        mt.total
        a.overall_avg
    }
}
```

### Ordering and Limiting
```preql
query {
    tables { product := p }
    order {
        p.sales desc
        p.name asc
    }
    limit 10
    offset 20
}
```

## Additional Ideas

1. **Type hints** for better IDE support:
```preql
tables {
    student<Student> := s
    department<Department> := d
}
```

2. **Computed fields** with explicit syntax:
```preql
get {
    s.email
    full_name := concat(s.first_name, ' ', s.last_name)
    is_honors := case when s.grade > 18 then true else false
}
```

3. **Explicit join types**:
```preql
joins {
    inner: s.department_id = d.id
    left: s.advisor_id = a.id
}
```

4. **Query composition**:
```preql
base_query students_in_math {
    tables { student := s, department := d }
    joins { s.department_id = d.id }
    conds { d.area = 'Math' }
}

query extends students_in_math {
    conds { s.grade > 16 }  // adds to existing conditions
    get { s.email, s.grade }
}
```

The key insight of preql - that SQL's backwards declaration order is problematic - is spot on. By reorganizing the syntax to match the mental model of query construction, you're making something that could be genuinely more pleasant to work with than SQL.
