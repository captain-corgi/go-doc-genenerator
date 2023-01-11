all: tidy vendor build run
build:
	go build -o demo cmd/demo/main.go
run:
	./demo
clean:
	rm ./demo
	rm -rf vendor

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy