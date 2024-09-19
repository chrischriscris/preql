# Grammar of preql

## General notes

- The expression terminator `<expr-terminator>` is technically a semicolon ( `;`),
    but the parser will insert a semicolon automatically on every newline if it
    makes sense to do so.

## Queries

```preql
<query-expression> -> query <query-block>
<query-block> ->
    <tables-expression>
    [<joins-expression>]
    [<conds-expression>]
    <get-expression>

<tables-expression> -> tables <tables-block>
<tables-block> -> {
    <table-decls>
}
<table-decls> -> (<table-decl>)+
<table-decl> -> <table-name> := <table-alias><expr-terminator>

<joins-expression> -> joins <joins-block>

<conds-expression> -> conds <conds-block>

<get-expression> -> get <get-block>
```
