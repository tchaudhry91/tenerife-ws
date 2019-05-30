
.PHONY: test
test:
	go test --race ./...

.PHONE: build
build:
	go build -o tenerife-ws ./cmd/tenerife-ws