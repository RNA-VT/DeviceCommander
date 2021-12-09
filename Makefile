lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

gqlgen:
	~/go/bin/gqlgen generate --verbose

mock:
	rm -rf mocks
	mockery --dir ./src --all --keeptree --output mocks

.PHONY: test
test:
	grc go test --cover ./...

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
