.PHONY: build
build:
	./scripts/build.sh

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix -v

gqlgen:
	~/go/bin/gqlgen generate --verbose

mock:
	rm -rf mocks
	./scripts/mock.sh

.PHONY: test
test:
	./scripts/test.sh

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
