# Backend Master Class (Golang + PostgreSQL + Kubernetes)

Application we will build: A simple bank

- Create and manage account: owner, balance, currency
- Record all balance changes: create an account entry for each change
- Money transfer transaction: perform money transfer between 2 accounts consistently within a transaction

## 1. Database design

- Design DB schema: design a SQL DB schema using `dbdiagram.io`
- Save & share DB diagram: save the DB schema as PDF/PNG diagram and share it with your team
- Generate code to create the schema in a target database engine: Postgres/MySQL/SQL Server

![DB](assets/db.png)

Code

```html
table accounts as A { id bigserial [pk] owner varchar [not null] balance bigint
[not null] currency varchar [not null] created_at timestamptz [default: `now()`]
// Search account by an owner indexes { owner } } // Record all changes to the
accounts table table entries { id bigserial [pk] account_id bigint [not null,
ref: > A.id] amount bigint [not null, note: 'can be positive or negative']
created_at timestamptz [default: `now()`] // List all entries of an account
indexes { account_id } } // Record all money transfer between two accounts table
transfer { id bigserial [pk] from_account_id bigint [not null, ref: > A.id]
to_account_id bigint [not null, ref: > A.id] amount bigint [not null, note:
'must be positive'] created_at timestamptz [default: `now()`] indexes {
from_account_id to_account_id (from_account_id, to_account_id) } }
```

## 2. Install Docker

![Image](assets/dockerimage.png)

### 2.1. Prepare Postgres Container

- Pulling the postgres image: `docker pull postgres:12-alpine`, using alpine version for smaller size.
- Run the container: `docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine`
- Run the command inside the container: `docker exec -it postgres12 psql -U root`

### 2.2. Connect to Postgres using TablePlus

- Open TablePlus
- Enter connected information
  - host: localhost
  - port: 5432
  - user: root
  - password: secret
  - database: root (default the same name as user name)

### 2.3. Create Table using TablePlus

- Export `dbdiagram.io` for `PostgreSQL`
- Open the downloaded file & copy all the content inside it
- Paste to the `SQL window` and execute those queries

## 3. Database Migration

Using `golang-migrate` library (`CLI tool`), [learn more](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

- Inside the project folder, create new folder for storing migration files: `mkdir -p db/migration`
- Check `golang-migrate` docs: `migrate -help`
- Create new migrate: `migrate create -ext sql -dir db/migration -seq init_schema`
- Paste the `sql` content to the `up` file
- Paste those line to the `down` file:

  ```sql
    DROP TABLE IF EXISTS entries;
    DROP TABLE IF EXISTS transfers;
    DROP TABLE IF EXISTS accounts;
  ```

- Create a `Makefile` for easily interact with the `container`

  ```Makefile
    postgres:
      docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

    createdb:
      docker exec -it postgres12 createdb --username=root --owner=root simple_bank

    dropdb:
      docker exec -it postgres12 dropdb simple_bank

    .PHONY: postgres createdb dropdb
  ```

- Use the `Makefile`:

  - Start new container: `make postgres`
  - Create new db: `make createdb`
  - Drop db: `make dropdb`

- Run the `migration`: `migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up`

## 4. CRUD

- Create: Insert new records to the database
- Read: Select or search for records in the database
- Update: Change some fields of the record in the database
- Delete: Remove records from the database

### 4.1. Implement CRUD in Golang

Things to consider

- DATABASE/SQL

  - Standard library
  - Very fast & straightforward
  - Manual mapping SQL fields to variables
  - Easy to make mistakes, not caught until runtime

- GORM

  - CRUD functions already implemented, very short production code
  - Must learn to write queries using gorm's function
  - Run slowly on high load

- SQLX

  - Quite fast & easy to use
  - Fields mapping via query text & struct tags
  - Failure won't occur until runtime

- SQLC

  - Very fast & easy to use
  - Automatic code generation
  - Catch SQL query errors before generating codes

### 4.2. CRUD with `sqlc`

- Intall: `yay -S sqlc`
- Init `sqlc` settings: `sqlc init`

```yaml
version: '1'
packages:
  - name: 'db'
    path: './db/sqlc'
    queries: './db/query/'
    schema: './db/migration/'
    engine: 'postgresql'
    emit_json_tags: true
    emit_prepared_queries: false
    emit_interface: true
    emit_exact_table_names: false
    emit_empty_slices: true
```

- Write raw query in `db/query`

```sql
-- name: CreateAccount :one
INSERT INTO accounts
    ( owner, balance, currency)
VALUES
    ( $1, $2, $3)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET
$3;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
```

- Generate code from raw query: `sqlc generate`

## 5. Unit Test

Create unit test for account

- Create `account_test.go` and place it in the same folder with the `account.sql.go`

Testing libraries:

- github.com/lib/pq v1.10.6 for `postgresql driver`
- github.com/stretchr/testify

### 5.1. Test entrypoint

Create `main_test.go` for the `test entrypoint`

- Setup `sql connection`
- Setup `testQueries`

```go
const (
 dbDriver = "postgres"
 dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
 conn, err := sql.Open(dbDriver, dbSource)

 if err != nil {
  log.Fatal("cannot connect to db:", err)
 }

 testQueries = New(conn)

 os.Exit(m.Run())
}
```

### 5.3. CRUD Test for account

```go
import (
 "context"
 "database/sql"
 "testing"
 "time"

 "github.com/dalatcoder/simple_bank/util"
 "github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
 arg := CreateAccountParams{
  Owner:    util.RandomOwner(),
  Balance:  util.RandomMoney(),
  Currency: util.RandomCurrency(),
 }

 account, err := testQueries.CreateAccount(context.Background(), arg)

 require.NoError(t, err)
 require.NotEmpty(t, account)

 require.Equal(t, arg.Owner, account.Owner)
 require.Equal(t, arg.Balance, account.Balance)
 require.Equal(t, arg.Currency, account.Currency)

 require.NotZero(t, account.ID)
 require.NotZero(t, account.CreatedAt)

 return account
}

func TestCreateAccount(t *testing.T) {
 createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
 account1 := createRandomAccount(t)
 account2, err := testQueries.GetAccount(context.Background(), account1.ID)

 require.NoError(t, err)
 require.NotEmpty(t, account2)

 require.Equal(t, account1.ID, account2.ID)
 require.Equal(t, account1.Owner, account2.Owner)
 require.Equal(t, account1.Balance, account2.Balance)
 require.Equal(t, account1.Currency, account2.Currency)

 require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
 account1 := createRandomAccount(t)

 arg := UpdateAccountParams{
  ID:      account1.ID,
  Balance: util.RandomMoney(),
 }

 account2, err := testQueries.UpdateAccount(context.Background(), arg)

 require.NoError(t, err)
 require.NotEmpty(t, account2)

 require.Equal(t, account1.ID, account2.ID)
 require.Equal(t, account1.Owner, account2.Owner)
 require.Equal(t, arg.Balance, account2.Balance)
 require.Equal(t, account1.Currency, account2.Currency)

 require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
 account1 := createRandomAccount(t)

 err := testQueries.DeleteAccount(context.Background(), account1.ID)
 require.NoError(t, err)

 account2, err := testQueries.GetAccount(context.Background(), account1.ID)
 require.Error(t, err)
 require.EqualError(t, err, sql.ErrNoRows.Error())
 require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
 for i := 0; i < 10; i++ {
  createRandomAccount(t)
 }

 arg := ListAccountsParams{
  Limit:  5,
  Offset: 5,
 }

 accounts, err := testQueries.ListAccounts(context.Background(), arg)

 require.NoError(t, err)
 require.Len(t, accounts, 5)

 for _, account := range accounts {
  require.NotEmpty(t, account)
 }
}
```

### 5.4. Run all tests

`go test -v -cover ./...`
