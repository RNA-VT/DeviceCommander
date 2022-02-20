.PHONY: build
build:
	./scripts/build.sh

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix -v

gqlgen:
	gqlgen generate --verbose

check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	swagger generate spec --work-dir src -o ./docs/swagger.json --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml

mock:
	rm -rf mocks
	./scripts/mock.sh

.PHONY: test
test:
	./scripts/test.sh

run:
	sudo go run main.go server
