GO_SRC=$(shell find . -type f -name '*.go')

.PHONY: run
run:
	air

.PHONY: build
build:
	go build -o build/app

.PHONY: fix
fix:
	@gopls format -w $(GO_SRC)
