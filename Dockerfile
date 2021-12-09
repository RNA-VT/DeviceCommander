FROM golang:1.17-alpine as go-builder

RUN apk add libpcap-dev build-base

WORKDIR /app

COPY . /app/

RUN go build -o device-commander

RUN chmod +x device-commander

# FROM node:15.11 as node-builder

# ADD /frontend /src

# WORKDIR /src

# RUN npm install

# RUN npm run build

# Build final image
FROM debian:10.4-slim

COPY --from=go-builder /app/device-commander /usr/local/bin/device-commander
# COPY --from=node-builder /src/build /src/build


ENV ENV "production"

EXPOSE 8000

