
build:
	cd src && go build
	
distribute:
	./environment/build-scripts/distribute-executables.sh

fix-permissions:
	chmod u+x ./environment/build-scripts/install-dependencies.sh
	chmod u+x ./environment/build-scripts/distribute-executables.sh

help:
	cd src && go run main.go -h

install:
	./environment/build-scripts/install-dependencies.sh

run-docker:
	cd environment && docker-compose up

run-master:
	cd src && GOFIRE_MASTER=true go run main.go

run-slave:
	cd src && go run main.go

run-slave2:
	cd src && GOFIRE_PORT=8002 GO111MODULE=on go run main.go
  
