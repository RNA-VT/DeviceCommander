#! /bin/bash

THISDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$THISDIR/config.sh"

pwd

for server in "${servers[@]}" ; do
    echo "[$server] Create Go/src Directory"
    sshpass -p $ssh_pass ssh pi@$server 'mkdir -p go/src'

    echo "[$server] Distribute executables"
    sshpass -p $ssh_pass scp firecontroller pi@$server:~/go/src/

    echo "[$server] Apply executable permissions"
    sshpass -p $ssh_pass ssh pi@$server 'sudo chmod +x ~/go/src/firecontroller'

    echo "[$server] Done distributing"
done