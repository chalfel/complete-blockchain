build:
	go build -o ./bin/complete-blockchain

run: build
	./bin/complete-blockchain

test:
	go test ./...