run:
	docker-compose up

run-old:
	go run main.go

run-master:
	go run main.go --makeMasterOnError

help:
	go run main.go -h

distribute:
	./build-scripts/distribute-executables.sh

fix-permissions:
	chmod u+x ./build-scripts/install-dependencies.sh
	chmod u+x ./build-scripts/distribute-executables.sh

install:
	./build-scripts/install-dependencies.sh


