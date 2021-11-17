lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

gqlgen:
	~/go/bin/gqlgen generate --verbose

mockery:
	mockery --dir ./src/postgres --all --keeptree --output src/mocks
	mockery --dir ./src/device --all --keeptree --output src/mocks

run-test:
	grc go test ./... -count=1

run:
	sudo go run main.go

# help:
# 	cd src && go run main.go -h

# build\:server:
# 	cd src && go build

# build\:client:
# 	cd frontend && npm run build

# run\:server:
# 	cd src &&  go run main.go

# run\:client:
# 	cd frontend && npm run start

# run\:docker:
# 	cd environment && docker-compose up
