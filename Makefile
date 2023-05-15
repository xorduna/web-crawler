lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o webcrawler main.go