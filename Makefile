.PHONY: test gen fmt

test:
	go test -v -cover ./...

gen:
	go generate

fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...
