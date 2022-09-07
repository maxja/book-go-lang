.DEFAULT_GOAL := build

fmt:
	go fmt exmpl_hello.go

lint: fmt
	golint exmpl_hello.go

vet: fmt
	go vet exmpl_hello.go

build: vet
	go build exmpl_hello.go

.PHONY:fmt lint vet build
