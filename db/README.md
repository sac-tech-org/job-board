# Database Migrations

The production db is hosted in a fly.io app, so to run a migration, we must do it locally by creating
a proxy connection to the fly.io Postgres app.

Create the proxy

```sh
fly proxy 5432 -a sac-tech-job-board-db
```

Then, in another terminal window

```sh
go run main.go
```
