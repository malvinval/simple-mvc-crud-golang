## Dependency

**`go get -u github.com/go-sql-driver/mysql`**

## DB Setup

- Create a new database with any name you want.
- Go to the `main.go` file, then set your database name as third parameter of `config.ConnectDB()` function.
- Import `db.sql` file provided in this repo to your database.
- Run command `go run main.go`
- Server will be deployed in `127.0.0.1:8000` (but you can customize the address from `main.go` file).