build:
	@go build -o bin/api-golang

run: build
	@./bin/api-golang

test: 
	@go test -v ./...