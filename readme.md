[![CircleCI](https://circleci.com/gh/RNA-VT/DeviceCommander/tree/master.svg?style=shield)](https://circleci.com/gh/RNA-VT/DeviceCommander/tree/master)

# Device Commander

Device Commander is the beginnings of an IOT platform. It will provide a useful layer for managing
a large collection of compliant devices.

## Getting Started

1. Build
```
  make build
```
2. Start Postgres
```
  make run-local-db
``` 
3. Run Migrations
```
  device-commander migrate-db
```
4. Start Device Commander
```
  device-commander server
```
