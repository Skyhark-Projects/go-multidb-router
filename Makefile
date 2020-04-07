# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

GOBIN = $(shell pwd)/build/bin
GO ?= latest

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

test:
ifdef bench
	build/env.sh go run build/ci.go test --bench=$(bench)
else
	build/env.sh go run build/ci.go test
endif

bench benchmark: all
ifdef fn
	build/env.sh go run build/ci.go test --bench=$(fn)
else
	build/env.sh go run build/ci.go test --bench=.
endif