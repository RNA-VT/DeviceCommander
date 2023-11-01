FROM golang:1.21-alpine as go-builder

ENV CGO_ENABLED=1

RUN apk update && apk add --update gcc libpcap-dev alpine-sdk

WORKDIR /app

COPY . /app/

RUN go build -o device-commander

RUN chmod +x device-commander

FROM node:alpine as node-builder

ADD /frontend /src

WORKDIR /src

RUN npm install

RUN npm run build

# Build final image
FROM debian:12-slim as final

COPY --from=go-builder /app/device-commander /usr/local/bin/device-commander
COPY --from=node-builder /src/build /src/build

ENV ENV "production"

EXPOSE 8000

