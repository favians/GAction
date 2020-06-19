all: build

build:
	go build -o actions

run:
	./actions

test:
	go test -v -cover -failfast
