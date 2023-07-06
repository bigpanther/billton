# warrant

## Install required tooling

Install the following on your machine.

- Go version [1.20+](https://go.dev/doc/install)
- XCode on your Mac from the App Store
- [flutter](https://docs.flutter.dev/get-started/install)

Add the following line in your `~/.zshrc`

```
export PATH=$PATH:$(go env GOPATH)/bin
```

Install the `soda` tool for database migrations

```
go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```

## run the program

To run the program, use

```
make run-dev
```

## Test Warranty Creation

From the terminal run

```
curl --location 'http://localhost:8080/warranty' \
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

```
curl --location 'http://localhost:8080/warranty/<replace this with the ID from the previous step>'
```
