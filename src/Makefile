lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

gqlgen:
	~/go/bin/gqlgen generate --verbose

mockery:
	mockery --dir ./postgres --all --keeptree

run-test:
	grc go test -v ./test -count=1

run:
	sudo go run main.go
