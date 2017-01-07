NAME=trainctl
ARCH=$(shell uname -m)
VERSION=0.0.3

build:
	mkdir -p build/Linux  && GOOS=linux  go build -ldflags "-X main.Version=$(VERSION)" -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/Darwin/$(NAME)

deps:
	go get -u -f github.com/jteeuwen/go-bindata/...
	go get || true

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/Darwin $(NAME)
	gh-release create gophertrain/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: release build

