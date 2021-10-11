.PHONY: fmt
fmt:
	gofmt -w .
	
.PHONY: build
build:
	go build -o nico main.go
