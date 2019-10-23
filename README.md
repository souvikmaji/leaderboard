# Leaderboard

Standalone REST web service to generate Fantasy sports leaderboard.

Leaderboards in MMOGs and Fantasy sports are the most accessed screens. What makes fantasy leaderboard little complex than any other rank-based leaderboards is that two or more players can share the same rank. The ranks of the players holding the same score should remain the same and the next succeeding rank to be allocated to the next player and so on.

In this project, we are using the SQL [rank function](http://www.sqltutorial.org/sql-window-functions/sql-rank/) to calculate team ranks.
In the SQL RANK functions, we use the `OVER()` clause to define the entire table is being considered for calculation of the rank. The `ORDER BY` clause is used to sort the results in descending order.

To see the database queries being generated while fetching data, set `database.logmode: true` in the config.yml file.

## Run in local environment:

### Prerequisites

1. Go version 1.13
2. make
3. PostgreSQL

Create a database from the psql console with `CREATE DATABASE databasename` or use the `createdb` tool (`createdb databasename`).

```sh
# Get the code
git clone https://github.com/souvikmaji/leaderboard
cd leaderboard

# DB setup from the database dump
psql databasename < scripts/dbdump

# Build and run
make
```

Visit: <127.0.0.1:8000>

## Technologies Used

- [Golang](https://golang.org/)
- [Gorilla Mux as the HTTP router and URL matcher](https://github.com/gorilla/mux)
- [Gorm as ORM](https://gorm.io/)
- [PostgreSQL as our database](https://www.postgresql.org/)
- [Datatables as the UI framework to display the leaderboard](https://datatables.net/)

## Things that can be improved / TODO:

- Unit tests
- Hosting
- Use datatable sorting feature
- refactor logic for query parameter parsing
- shields
- test with a larger dataset
