# Leaderboard

![GitHub](https://img.shields.io/github/license/souvikmaji/leaderboard) [![Maintainability](https://api.codeclimate.com/v1/badges/2a9a5bb2bf0935181799/maintainability)](https://codeclimate.com/github/souvikmaji/leaderboard/maintainability)

Standalone REST web service to generate Fantasy sports leaderboard.

Leaderboards in MMOGs and Fantasy sports are the most accessed screens. What makes fantasy leaderboard little complex than any other rank-based leaderboards is that two or more players can share the same rank. The ranks of the players holding the same score should remain the same and the next succeeding rank to be allocated to the next player and so on.

In this project, we are using the SQL [rank function](http://www.sqltutorial.org/sql-window-functions/sql-rank/) to calculate user ranks.

The `OVER()` clause In the SQL RANK function is used to define that the entire table is being considered for calculation of the rank. The `ORDER BY` clause is used to sort the results in descending order.

To see the database queries being generated while fetching data, set `database.logmode: true` in the config.yml file.

## Run in local environment

### Prerequisites

1. Go with minimum version 1.13
2. make
3. PostgreSQL

### Setup

```sh
# Get the code
git clone https://github.com/souvikmaji/leaderboard
cd leaderboard

# initialize database and tables
make init

# Build and run
make
```

The application follows [12 factor](https://12factor.net/config) principals for configuration management. Configurations are read in the following manner `config.yml > .env > Exported Enviroment variables > Program Flags`.

Visit: <127.0.0.1:8000>

### Live Reload Server

During development use the autoreload server

```sh
make devrun
```

## Technologies Used

- [Golang](https://golang.org/)
- [GNU Make](https://www.gnu.org/software/make/manual/html_node/Introduction.html)
- [Gorilla Mux as the HTTP router and URL matcher](https://github.com/gorilla/mux)
- [Gorm as ORM](https://gorm.io/)
- [Viper for externalized configuration](https://github.com/spf13/viper)
- [PostgreSQL as the database](https://www.postgresql.org/)
- [Datatables as the UI framework to display the leaderboard](https://datatables.net/)

# Things that can be improved / TODO

- CRUD apis for game
- CRUD apis for users
- Use datatable sorting feature
- Handle schema decoder errors
- Refactor setup router logic from a map
- User auth apis
- Gameplay apis
- DB script sample size from flag
- Use redis for db
- Benchmark with a larger dataset
- Read database credentials at app start for the first time.
- Integrate Swagger
- AutomaticEnv not working
- Add logger
- Hosting
- Unit tests
