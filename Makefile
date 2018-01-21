GOPATHDIR = gopath
PROJNAME = github.com/rwdalpe/packer-builder-aws-stopped-instance-pool
GOPROJSRCPATH = $(GOPATHDIR)/src/$(PROJNAME)

export GOPATH := $(abspath $(GOPATHDIR))

copy_src_to_gopath:
	mkdir -p "$(GOPROJSRCPATH)"
	find . -maxdepth 1 \
		-not -name ".*" \
		-not -name "$(GOPATHDIR)" \
		-not -name "Makefile" \
		-not -name ".gitignore" \
		-exec cp -R {} "$(GOPROJSRCPATH)" \;

fmt:
	find . \
		-not -path "./$(GOPATHDIR)/*" \
		-name "*.go" \
		-exec go fmt {} \;

get: copy_src_to_gopath
	cd "$(GOPROJSRCPATH)" && go get ./...

test: clean fmt get
	cd "$(GOPROJSRCPATH)" && go test ./...

build: clean fmt get
	go build $(PROJNAME)

install: clean fmt get
	go install $(PROJNAME)

clean:
	rm -rf $(GOPROJSRCPATH)

clean-go:
	rm -rf $(GOPATHDIR)
