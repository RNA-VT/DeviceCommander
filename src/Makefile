lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

gqlgen:
	~/go/bin/gqlgen generate --verbose

mockery:
	mockery --dir ./postgres --all --keeptree

test:
	go test ./test -count=1 -v

run:
	sudo go run main.go
