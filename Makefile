.PHONY: build
build:
	go build

.PHONY: test
test: build
	./travis-ci-experiments
