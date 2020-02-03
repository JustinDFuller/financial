.PHONY: test

test:
	goimports -w ./*.go;
	gofmt -s -w ./*.go;
	go test -race ./...;

