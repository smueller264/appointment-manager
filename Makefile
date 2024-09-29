build: 
	@go build -o bin/appointment-manager

run: build 
	@./bin/appointment-manager

test: 
	@go test -v ./...