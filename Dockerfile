FROM golang:1.13.5 as go-builder

RUN mkdir -p /go/src/DeviceCommander/

ADD /src /go/src/DeviceCommander/

WORKDIR /go/src/DeviceCommander/

RUN go build -o device-commander

RUN ls

RUN chmod +x device-commander

FROM node:15.11 as node-builder

ADD /frontend /src

WORKDIR /src

RUN npm install

RUN npm run build

# Build final image
FROM debian:10.4-slim

COPY --from=go-builder /go/src/DeviceCommander/device-commander /usr/local/bin/device-commander
COPY --from=node-builder /src/build /src/build

RUN touch /config.yaml

ENV ENV "production"

EXPOSE 8000

