# Vara_backend_go

 A video site demo with go.

## todo list

* [x] Graceful shutdown
* [x] Recover
* [x] Logger
* [x] Config
* [x] Database
* [x] Migration
* [x] User
* [x] JWT
* [ ] Video
* [ ] Image
* [ ] Forum

## packages

* github.com/labstack/echo/v4
* github.com/labstack/echo-jwt/v4
* github.com/jackc/pgx/v5/pgxpool
* github.com/vgarvardt/pgx-google-uuid/v5
* github.com/google/uuid
* github.com/spf13/viper
* go.uber.org/zap
* golang.org/x/crypto

## migrate

```bash
goose postgres "user=postgres password=password dbname=vara sslmode=disable" status
goose postgres "postgres://postgres:password@localhost:5433/vara?sslmode=disable" status
```

```
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
    validate             Check migration files without running them
```

### seed

apply all migrations 

```bash
goose -dir ./seed -no-versioning postgres "user=postgres password=password dbname=vara sslmode=disable" up
```

apply all down migrations (same as down-to 0)
```bash
goose -dir ./seed -no-versioning postgres "user=postgres password=password dbname=vara sslmode=disable" reset
```