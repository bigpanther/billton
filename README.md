# warrant

## Install required tooling

Install the following on your machine.

- Go version [1.20+](https://go.dev/doc/install)
- XCode on your Mac from the App Store
- [Flutter](https://docs.flutter.dev/get-started/install) fpr mobile development
- [Postman](https://www.postman.com/downloads/) for API access

Add the following line in your `~/.zshrc`. In the terminal, run the following command

```bash
grep 'export PATH=$PATH:$(go env GOPATH)/bin' ~/.zshrc || echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
```

Install the `soda` tool for database migrations

```bash
go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

## run the program

To run the program, use

```bash
make run-dev
```

## Test Warranty Creation

Open a new terminal and run

```bash
curl --location 'http://localhost:8080/warranties' \
--header 'Content-Type: application/json' \
--data '{
    "transaction_time": "2023-06-29T20:52:20.015924-07:00",
    "expiry_time": "2024-07-03T20:52:20.015924-07:00",
    "brand_name": "Samsung",
    "amount": 100000,
    "store_name": "Costco"
}'
```

Expect to get back a response with a generated ID

### Test warranty fetch

```bash
curl --location 'http://localhost:8080/warranties/<replace this with the ID from the previous step>'
```

## Add a new model and a database table

```
soda generate fizz create_users
```

In the migration up file, add the new table schema

```fizz
create_table("users") {
    t.Column("id", "uuid", {primary: true})
    t.Column("name", "string", {})
    t.Timestamps()
}

```

In the down migration file, add the reverse

```fizz
drop_table("users")
```

See [db migrations](https://gobuffalo.io/documentation/database/migrations/) and [fizz](https://gobuffalo.io/documentation/database/fizz/) for details on migrations.

Create a model under the `models` directory to map the DB table

```go
package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id" rw:"r"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

```
