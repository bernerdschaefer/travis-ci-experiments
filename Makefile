.PHONY: build
build:
	go build

.PHONY: test
test: build
	./travis-ci-experiments

.PHONY: test-sudo
test-sudo: build
	sudo ./travis-ci-experiments
