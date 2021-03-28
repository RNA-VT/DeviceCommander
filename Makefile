help:
	cd src && go run main.go -h

build\:server:
	cd src && go build

build\:client:
	cd frontend && npm run build

run\:server:
	cd src &&  go run main.go

run\:client:
	cd frontend && npm run start

run\:docker:
	cd environment && docker-compose up
