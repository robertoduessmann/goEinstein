# goEinstein

A Go API to calculate the famous equation E=mc²

## Build
```sh
go build
```
## Run
```sh
./goEinstein
```
## Usage
```sh
curl http://localhost:3000/einsten/{mass}
```
## Example
#### Request
```sh
curl http://localhost:3000/einsten/5
```
#### Response
```sh
450000000000000000.000000
```