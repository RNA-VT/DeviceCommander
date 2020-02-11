## GoFire 

## Dependencies

  - sshpass
    -  `brew install http://git.io/sshpass.rb` (MAC)

## Environment Variables
  GOFIRE_HOST node hostname or ip (Default: 127.0.0.1)
  GOFIRE_PORT node service port (Default: 8001)
  GOFIRE_MASTER Set to true to let this node take on the master role if a master cannot be reached (Default: false)
  GOFIRE_MASTER_HOST Master service hostname (Default: 127.0.0.1) 
  GOFIRE_MASTER_PORT Master service port (Default: 8000)
	GOFIRE_MOCK_GPIO Activates mock mode for all io modules (Default: true)
	MICROCONTORLLER_LIMIT sets the maximum number of devices this cluster can accept (Default: 255)
  

## Example Start
  `GOFIRE_HOST=127.0.0.1 GOFIRE_PORT=8001 GOFIRE_MASTER=false go run main.go`

## Make Commands

  build: build Go source
  
  distribute: manually build and deploy GoFire to the ips in config.sh

  fix-permissions: permission install scrips

  help: show GoFire help

  install: install GoFire deps

  run-docker: run GoFire locally in a docker container

  run-master: run GoFire in Master mode
  
  run-slave: run GoFire in Slave mode (port 8001)

  run-slave2: run GoFire in Slave mode (port 8002)

## Tips
  Make sure you have your go modules enabled: `export GO111MODULE=on`
  
