GOCACHE := $(CURDIR)/.gocache

.PHONY: serve build air

serve:
	env GOCACHE=$(GOCACHE) go run ./cmd/gobbs-serve

build:
	env GOCACHE=$(GOCACHE) go run ./cmd/gobbs-static -out public

air:
	air
