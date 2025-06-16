build:
	@go build -o bin/blkchain

run: build
	@./bin/blkchain

test:
	@go test ./... -v