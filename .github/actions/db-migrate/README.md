# DB Migrate docker action

This action runs the `migrate up` command from [golang-migrate](https://github.com/golang-migrate/migrate).
It does not take any arguments, it is intended to be simple. It requires an environment variable
POSTGRES_URL to be set to be able to connect to the database.
