## GoFire 

## Dependencies

  - sshpass
    -  `brew install http://git.io/sshpass.rb` (MAC)

## Environment Variables
  GOFIRE_HOST Master node hostname or ip (Default: 127.0.0.1)
  GOFIRE_PORT Master node service port (Default: 8001)
  GOFIRE_MASTER Set to true to let this node take on the master role if a master cannot be reached (Default: false)

## Example Start
  `GOFIRE_HOST=127.0.0.1 GOFIRE_PORT=8001 GOFIRE_MASTER=false ENV=local go run main.go`

## Make Commands

  build: build Go source
  
  distribute: manually build and deploy GoFire to the ips in config.sh

  fix-permissions: permission install scrips

  help: show GoFire help

  install: install GoFire deps

  run-docker: run GoFire locally in a docker container

  run-master: run GoFire in Master mode

  run-slave: run GoFire in Slave mode