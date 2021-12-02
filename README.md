# demos-go

## Run
go run .

## Test
go test -v

## Build
Generate executable
go build
Example:
./hello

## Install
Install executable
go install

## Create module
go mod init example.com/hello

## Redirect to local module (hello directory)
go mod edit -replace example.com/greetings=../greetings

## Synchronize the example.com/hello module's dependencies
go mod tidy

