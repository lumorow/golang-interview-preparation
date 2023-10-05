# PostgreSQL

In PostgreSQL, a transaction is defined by a set of SQL commands surrounded by `BEGIN` and `COMMIT` commands.
So our bank transaction would look like this:

```sql
BEGIN;
UPDATE accounts SET balance = balance - 100.00
WHERE name = 'Alice';
-- ...
COMMIT;
```
If during the execution of a transaction we decide that we do not want to commit its changes (for example, because it turned out that
that Alice's balance has become negative), we can issue the command `ROLLBACK` instead of `COMMIT`, and all our changes will be
cancelled.

PostgreSQL actually processes each SQL statement as a transaction. If you do not insert the `BEGIN` command,
then each individual statement will be implicitly surrounded by `BEGIN` and `COMMIT` commands (if successful).
A group of statements surrounded by `BEGIN` and `COMMIT` commands is sometimes called a transaction block.

## ACID or 4 transaction properties

Before we begin to consider transaction isolation levels, let us briefly recall the basic requirements for a transactional system.

- `Atomicity` - is expressed in the fact that the transaction must be completed as a whole or not performed at all.

- `Consistency` - ensures that as transactions are performed, data moves from one consistent state
  to another, that is, the transaction cannot destroy the mutual consistency of the data.

- `Isolation` - localization of user processes means that transactions competing for access to the database are physically processed sequentially,
  isolated from each other, but to users it appears as if they are running in parallel.

- `Durability`- error resistance - if a transaction is completed successfully, then the changes in the data that were made by it cannot be lost under any circumstances.

## Transactions and isolation levels

The standard describes the following special conditions that are not permitted for the various insulation levels:

- `Dirty reading`. A transaction reads data written by a concurrent pending transaction.

- `Unrepeatable reading`. The transaction re-reads the same data as before and discovers that it was modified by another transaction (which completed after the first read).

- `Phantom reading`. A transaction re-executes a query that returns a set of rows for some condition and discovers that the set of rows that satisfy the condition has changed due to a transaction that completed during that time.

- `Serialization anomaly`. The result of successfully committing a group of transactions turns out to be inconsistent under all possible options for executing these transactions in turn.

| Isolation level   | Dirty r eading         | Unique reading | Phantom reading        | Serialization anomaly |
|-------------------|------------------------|----------------|------------------------|-----------------------|
| `Read uncommited` | Allowed, but not in PG | Perhaps        | Perhaps                | Perhaps               |
| `Read committed`  | Impossible             | Perhaps        | Perhaps                | Perhaps               |
| `Repeatable read` | Impossible             | Impossible     | Allowed, but not in PG | Perhaps               |
| `Serializable`    | Impossible             | Impossible     | Impossible             | Impossible            |

## Indexes in PostgreSQL

To improve search performance, auxiliary structures - indexes - are created. Using indexes,
we can significantly increase the search speed, because the data in the index is stored in a form that allows us to
search does not consider areas that obviously cannot contain the searched elements.

PostgreSQL supports different types of indexes for different tasks:

- `B-tree` covers the widest class of problems, since it is applicable to any data that can be sorted.
- `GiST` and `SP-GiST` can be useful when working with geometric objects and for creating completely new types of indexes
  for new data types.
- Beyond the scope of this article was another important type of index - `GIN`. `GIN` indexes are useful for organizing full-text
  search and for indexing data types such as arrays or `jsonb`.


## Resources

- [Transaction isolation in PostgreSQL (DEV Meetup November 2019)](https://www.youtube.com/watch?v=h4-o0kQz5lo&t=11s&ab_channel=QAHub)
- [Transaction isolation (PostgresPro)](https://postgrespro.ru/docs/postgrespro/9.5/transaction-iso)
- [Transaction isolation levels with examples in PostgreSQL (habr)](https://habr.com/ru/articles/317884/)
- [Indexes in PostgreSQL - 1 (habr)](https://habr.com/ru/companies/postgrespro/articles/326096/)
- [Indexes in PostgreSQL](https://tproger.ru/articles/indeksy-v-postgresql)
- [Indexes in PostgreSQL - 3 (habr)](https://habr.com/ru/companies/postgrespro/articles/328280/)

## README.md

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/PosgreSQL/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/PosgreSQL/readme/README.ru.md)
