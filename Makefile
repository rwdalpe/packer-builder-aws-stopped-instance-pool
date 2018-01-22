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
		-not \( -path "./$(GOPATHDIR)" -prune \) \
		-name "*.go" \
		-exec go fmt {} \;

vendor: copy_src_to_gopath
	cd "$(GOPROJSRCPATH)" && govendor sync

test: clean fmt vendor
	cd "$(GOPROJSRCPATH)" && go test ./...

build: clean fmt vendor
	go build $(PROJNAME)

install: clean fmt vendor
	go install $(PROJNAME)

clean:
	[[ -d "$(GOPROJSRCPATH)" ]] && \
	find $(GOPROJSRCPATH) -maxdepth 1 \
		-not -name ".*" \
		-not -name "vendor"
		-exec rm -rf {} \;

clean-vendor:
	rm -rf "$(GOPROJSRCPATH)/vendor"

clean-go:
	rm -rf $(GOPATHDIR)
