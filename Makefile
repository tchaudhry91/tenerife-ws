
.PHONY: test
test:
	go test --race ./...

.PHONY: build
build:
	go build -o tenerife-ws ./cmd/tenerife-ws