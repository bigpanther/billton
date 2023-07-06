run-dev: migrate-dev
	GO_ENV=dev go run -tags sqlite  main.go

migrate-dev:
	GO_ENV=dev soda migrate

migrate-test:
	GO_ENV=test soda migrate

test: migrate-test
	GO_ENV=dev go test -tags sqlite ./... -v

drop-dev:
	GO_ENV=dev soda drop
drop-test:
	GO_ENV=test soda drop
