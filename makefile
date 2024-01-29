run:
	go run *.go

build:
	GOARCH=amd64 GOOS=linux go build *.go