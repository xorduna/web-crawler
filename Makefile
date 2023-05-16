lint:
	golangci-lint run

test:
	go test -coverprofile=profile.cov ./...
	go tool cover -func profile.cov
	rm profile.cov

build:
	go build -o webcrawler cmd/main.go