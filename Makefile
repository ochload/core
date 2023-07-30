GO_TEST_CMD=go test
GO_TEST_PATHS=./...
GO_GENERATE=go generate
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := $(LDFLAGS) -X main.commit=$(COMMIT)
VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
ifdef VERSION
	LDFLAGS += -X main.version=$(VERSION)
endif

export GO_TEST=env GOTRACEBACK=all GO111MODULE=on $(GO_TEST_CMD)
export GO_BUILD=env GO111MODULE=on go build -ldflags "$(LDFLAGS)"
export GO_RUN=env GO111MODULE=on go run -ldflags "$(LDFLAGS)"
export NPX_BUILD=npx webpack

test: 
	$(GO_TEST) $(GO_TEST_PATHS)

build:
	$(NPX_BUILD)
	$(GO_BUILD) -o ./bin/ochload ./cmd/ochload/.

run:
	$(NPX_BUILD)
	$(GO_RUN) ./cmd/ochload/.