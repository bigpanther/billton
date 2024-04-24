# warrant

## Install required tooling

Install the following on your machine.

- Go version [1.21+](https://go.dev/doc/install)
- Postgres [16.1+](https://www.postgresql.org/download/macosx/)
- XCode on your Mac from the App Store
- [Flutter](https://docs.flutter.dev/get-started/install) fpr mobile development
- [Postman](https://www.postman.com/downloads/) for API access

Add the following line in your `~/.zshrc`. In the terminal, run the following command

```bash
grep 'export PATH=$PATH:$(go env GOPATH)/bin' ~/.zshrc || echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
```

Clone this repository

```
git clone https://github.com/bigpanther/billton.git
```

Open the code in `Visual Studio Code`

Install the `atlas` tool for database migrations

```bash
brew install ariga/tap/atlas
```

## Fetching latest code from Github (Terminal)

```
git checkout main
git pull --rebase

git checkout <your-branch>
git rebase main

# Make changes and commit the code

# Push the changes to github
git push origin <your-branch>
```

You can use the VSCode git extension to run the same steps from the UI instead of the terminal.

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

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id" rw:"r"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

```
To use WarranT app, follow these steps:

1. Clone the following repository to your local machine using the following code.
```
git clone https://github.com/bigpanther/billton.git
```
2. Ensure you have Flutter installed on your system.
3. Run `flutter pub get` to install dependencies.
4. Ensure to open simualator when testing the app on your computer
5. Run the app using `flutter run`.

## Dependencies

The WarranT app uses the following dependencies:
  - cupertino_icons: ^1.0.2
  - firebase_core: ^2.26.0
  - firebase_auth: ^4.17.7
  - google_sign_in: ^6.2.1
  - form_builder_validators: ^9.1.0
  - flutter_form_builder: ^9.2.1
  - image_picker: ^0.8.9 
  - firebase_storage: ^11.6.8

## Contributing

Contributions are welcome! If you encounter any bugs or have suggestions for new features, feel free to open an issue or submit a pull request.

## Contact

If you have any questions or need further assistance, you can contact the developer:

Harsimran Singh Maan at harsimran.maan@indcan.org

