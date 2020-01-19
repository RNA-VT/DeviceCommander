FROM golang:1.13.5

RUN mkdir -p /go/src/GoFire/

RUN export GOFIRE_MASTER_HOST=`/sbin/ip route|awk '/default/ { print $3 }'` && export GOFIRE_MASTER=true

ADD /src /go/src/GoFire/

WORKDIR /go/src/GoFire/

RUN go build

RUN chmod +x firecontroller

RUN GOFIRE_MASTER=true ./firecontroller
