#! /bin/bash

THISDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$THISDIR/config.sh"

pwd

for server in "${servers[@]}" ; do
    if sshpass -p $ssh_pass test -s ~/https://dl.google.com/go/go1.13.4.linux-armv6l.tar.gz
    then
        echo "[$server] WGET Go"
        sshpass -p $ssh_pass ssh pi@$server 'wget https://dl.google.com/go/go1.13.4.linux-armv6l.tar.gz'

        echo "[$server] Un-Tar Go"
        sshpass -p $ssh_pass ssh pi@$server 'sudo tar -C /usr/local -xzf go1.13.4.linux-armv6l.tar.gz'

        echo "[$server] Export Go to PATH"
        sshpass -p $ssh_pass ssh pi@$server 'echo "export GOPATH=$HOME/go" >> ~/.profile'
        sshpass -p $ssh_pass ssh pi@$server 'echo "export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin" >> ~/.profile'
        sshpass -p $ssh_pass ssh pi@$server 'source ~/.profile'
    fi

    echo "[$server] Done installing"
done