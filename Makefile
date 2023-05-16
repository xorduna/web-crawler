lint:
	golangci-lint run

test:
	go test -v ./...

build:
	go build -o webcrawler cmd/main.go