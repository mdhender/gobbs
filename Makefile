GOCACHE := $(CURDIR)/.gocache
BROWSER_SYNC ?= npx --yes browser-sync
PREVIEW_ADDR ?= 127.0.0.1:8080

.PHONY: serve build air live

serve:
	env GOCACHE=$(GOCACHE) go run ./cmd/gobbs-serve

build:
	env GOCACHE=$(GOCACHE) go run ./cmd/gobbs-static -out public

air:
	air

live:
	@trap 'kill 0' EXIT INT TERM; \
	$(MAKE) air & \
	BROWSER_SYNC_PROXY=http://$(PREVIEW_ADDR) $(BROWSER_SYNC) start --config browser-sync.config.js
