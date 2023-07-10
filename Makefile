.PHONY: build
build:
	./build.sh

.PHONY: run
run:
	./output/bootstrap.sh

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: build-image
build-image:
	docker build -t hertz_app .

.PHONY: clean
clean:
	rm -rf output