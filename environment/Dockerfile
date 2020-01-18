FROM golang:1.13.5

RUN mkdir -p /go/src/GoFire

RUN export GOFIRE_MASTER_HOST=`/sbin/ip route|awk '/default/ { print $3 }'` && export GOFIRE_MASTER=true

ADD . /go/src/GoFire

WORKDIR /go/src/GoFire/src

RUN ls
