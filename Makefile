.PHONY: run
run:
	air

.PHONY: build
build:
	go build -o build/app

.PHONY: fix
fix:
	go fmt