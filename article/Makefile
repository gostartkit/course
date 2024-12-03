GIT_COMMIT := $(shell git rev-parse HEAD)

RELEASES=bin/article-darwin-amd64 \
	bin/article-darwin-arm64 \
	bin/article-windows-amd64.exe \
	bin/article-linux-amd64

all: $(RELEASES)

bin/article-%: GOOS=$(firstword $(subst -, ,$*))
bin/article-%: GOARCH=$(subst .exe,,$(word 2,$(subst -, ,$*)))
bin/article-%: $(wildcard *.go)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build \
	     -ldflags "-X 'main.osarch=$(GOOS)/$(GOARCH)' -X 'main.gitrev=$(GIT_COMMIT)' -s -w" \
	     -buildmode=exe \
	     -tags release \
	     -o $@

clean:
	rm -rf bin
