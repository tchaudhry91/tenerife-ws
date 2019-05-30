
.PHONY: test
test:
	go test --race ./...

.PHONY: build
build:
	go build -o tenerife-ws ./cmd/tenerife-ws

.PHONY: clean
clean:
	rm tenerife-ws || true
	docker rmi tenerife-ws || true

.PHONY: docker
docker:
	docker build -t tenerife-ws .