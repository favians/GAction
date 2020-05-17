all: build

run:
	go run main.go hello.go

test:
	go test -v -cover -failfast
