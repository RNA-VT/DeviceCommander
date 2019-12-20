go run main.go

## Dependencies

  - sshpass
    -  `brew install http://git.io/sshpass.rb` (MAC)

## Environment Variables
  HOST Master node hostname or ip (Default: 127.0.0.1)
  PORT Master node service port (Default: 8001)
  DEFAULT_MASTER Set to true to let this node take on the master role if a master cannot be reached (Default: false)

## Example Start
  `HOST=127.0.0.1 PORT=8001 DEFAULT_MASTER=false ENV=local go run main.php`
