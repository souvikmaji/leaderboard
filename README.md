# Leaderboard

![GitHub](https://img.shields.io/github/license/souvikmaji/leaderboard) ![Codacy grade](https://img.shields.io/codacy/grade/a0b36eaed7b74ec6be460ceda499bf2c)

Standalone REST web service to generate Fantasy sports leaderboard.

Leaderboards in MMOGs and Fantasy sports are the most accessed screens. What makes fantasy leaderboard little complex than any other rank-based leaderboards is that two or more players can share the same rank. The ranks of the players holding the same score should remain the same and the next succeeding rank to be allocated to the next player and so on.

In this project, we are using the SQL [rank function](http://www.sqltutorial.org/sql-window-functions/sql-rank/) to calculate team ranks.

The `OVER()` clause In the SQL RANK function is used to define that the entire table is being considered for calculation of the rank. The `ORDER BY` clause is used to sort the results in descending order.

To see the database queries being generated while fetching data, set `database.logmode: true` in the config.yml file.

## Run in local environment

### Prerequisites

1.  Go with minimum version 1.13
2.  make
3.  PostgreSQL

### Setup

Create a database from the psql console with `CREATE DATABASE leaderboard;` or use the `createdb` tool (`createdb leaderboard`).

Check if the database is being created successfully using `\l`. It will list all the databases for the current user.

Run the following script to complete setup.

```sh
# Get the code
git clone https://github.com/souvikmaji/leaderboard
cd leaderboard

# Initialize data from the database dump
psql leaderboard < scripts/dbdump

# Build and run
make
```

The default db name is `leaderboard`. To change the database name update the `config.yml` file or export corresponding environment variables.

Update

For running psql from a postgres user add `sudo -u postgres` at the begining of the command.

Visit: [127.0.0.1:8000](127.0.0.1:8000)

## Technologies Used

-   [Golang](https://golang.org/)
-   [GNU Make](https://www.gnu.org/software/make/manual/html_node/Introduction.html)
-   [Gorilla Mux as the HTTP router and URL matcher](https://github.com/gorilla/mux)
-   [Gorm as ORM](https://gorm.io/)
-   [Viper for externalized configuration](https://github.com/spf13/viper)
-   [PostgreSQL as the database](https://www.postgresql.org/)
-   [Datatables as the UI framework to display the leaderboard](https://datatables.net/)

## Things that can be improved / TODO

-   Unit tests
-   Hosting
-   Use datatable sorting feature
-   Read database credentials at app start for the first time.
-   Test with a larger dataset
-   Initialize db using make
-   Generalize team/player nomanclatures.
-   Handle schema decoder errors
-   Refactor setup router logic into a map in controller package
