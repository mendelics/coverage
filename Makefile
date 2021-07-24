.PHONY: build
build:
	@echo "\033[0;33mbuilding coverage binaries\033[m"

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/coverage-linux-amd64 github.com/mendelics/coverage

	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/coverage-linux-arm64 github.com/mendelics/coverage

	GOOS=darwin GOARCH=amd64 go build -o bin/coverage-macos-amd64 github.com/mendelics/coverage

	GOOS=darwin GOARCH=arm64 go build -o bin/coverage-macos-arm64 github.com/mendelics/coverage

	@echo "\033[0;33mcoverage installed\033[m"
