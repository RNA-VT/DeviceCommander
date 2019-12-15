go run main.go

## Dependencies

  - sshpass
    -  `brew install http://git.io/sshpass.rb` (MAC)

## Run Args
  -clusterip string
    	ip address of any node to connnect (default "127.0.0.1:8001")
  -makeMasterOnError
    	make this node master if unable to connect to the cluster ip provided.
  -myport string
    	ip address to run this node on. default is 8001. (default "8001")

## Example Start
  `go run main.php --myport 8002 --clusterip 8001`
