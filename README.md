# Leaderboard

MMOG Leaderboard

## Run in local environment:

### Prerequisites

1. Go version 1.13
2. make
3. PostgreSQL

Create a database from the psql console with `CREATE DATABASE databasename` or use the `createdb` tool.

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
