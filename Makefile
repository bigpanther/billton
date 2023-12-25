run-dev: migrate-dev
	GO_ENV=dev go run main.go

migrate-dev:
	atlas schema apply --env dev

migrate-test:
	atlas schema apply --env test

test: migrate-test
	GO_ENV=test go test ./... -v

gen-migration:
	atlas migrate diff --env dev

drop-dev:
	atlas schema clean --env dev
drop-test:
	atlas schema clean --env test

build:
	GO_ENV=dev go build -o release/warrant -v .
