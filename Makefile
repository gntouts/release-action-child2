COMMIT         := $(shell git describe --dirty --long --always)
VERSION        := $(shell cat $(CURDIR)/VERSION)-$(COMMIT)
LDFLAGS_COMMON := -X main.version=$(VERSION)

build:
	CGO_ENABLED=0 go build -a -ldflags="$(LDFLAGS_COMMON) -s -w -extldflags=-static" -trimpath -o $(CURDIR)/dist/reversecho main.go