
build: fmt
	go build

fmt:
	find . -name '*.go' -exec gofmt -w -s {} \;