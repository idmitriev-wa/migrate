# Migrate

Purpose: automation of the migration to the database using the CLI.

Supported databases:

- PostgreSQL
- Cassandra

Base library: https://github.com/mattes/migrate

## Installation

To install the library and command line program, use the following:
```
$ go get github.com/idmitriev-wa/migrate/
```

To add dependencies, use the following:
```
go get $(<requirements.txt)
```

## Usage

Import migrate package into your application:
```
 import github.com/idmitriev-wa/migrate/
```

Set up a source of migrations as path string (in example migrations are located in **/migrations/postgres/** folder) and database connection string:
```
migrate.ExecMigrate("/migrations/postgres/", "postgres://user:password@host:5433/dbname")
```

Check whether the migration command:
```
if !migrate.IsExecMigrate() {
    panic("Incorrect command")
}
```

#### Run from CLI:

Get version your current migration:
```
go run *.go migrate version
```

Create new migration:
```
go run *.go migrate create migrate_name
```

Apply all migration:
```
go run *.go migrate up
```

Apply the first n migration:
```
go run *.go migrate up n
```

Cancel all migration:
```
go run *.go migrate down
```

Cancel the first n migration:
```
go run *.go migrate down n
```

Restart the latest migration:
```
go run *.go migrate redo
```