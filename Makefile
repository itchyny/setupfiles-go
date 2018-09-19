all: lint test

test:
	go test -v ./...

lint: lintdeps
	golint -set_exit_status *.go

lintdeps:
	command -v golint >/dev/null || go get -u github.com/golang/lint/golint

.PHONY: test lint lintdeps
