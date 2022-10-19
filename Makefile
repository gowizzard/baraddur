# Here you can reformat, check or build the binary.

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test ./...

lint:
	@golangci-lint run ./...

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060