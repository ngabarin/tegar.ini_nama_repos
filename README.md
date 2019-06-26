# Golang Boilerplate for roadrunner

## Install dependencies

go install goBoilerplate/main.go

## Run the test

go test goBoilerplate/...

## Run the app

go run goBoilerplate/main.go

## Build the app

go build goBoilerplate/main.go

# Docker 

Build:

```
docker build -t tegarimansyah/hello-go-multistage .
```

Run

```
docker run --rm -p 1323:1323 tegarimansyah/hello-go-multistage
```