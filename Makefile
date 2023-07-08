.PHONY: test test-cover gen gen-ic fmt

test:
	go test -v -cover ./...

gen:
	go generate

fmt:
	go mod tidy
	gofmt -s -w .
	goarrange run -r .
	golangci-lint run ./...
