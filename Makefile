PRODUCT_VERSION          ?= 1
PRODUCT_REVISION         ?= 2
BUILD_VERSION            ?= $(PRODUCT_VERSION)-$(PRODUCT_REVISION)

clean:
	echo "Cleaning the project"
	go clean
	rm -rf ./bin

build-install: clean
	$(info) "Building version: $(BUILD_VERSION)"
	GOBIN=$$PWD/bin/ go install -ldflags="$(LDFLAGS)" $(get_packages)
	golint -set_exit_status $(get_packages)
	go vet $(get_packages)

test: build
	go test $(get_packages)

build: build-install
get_packages := $$(go list ./... | grep -v /vendor/)
info := @printf "\033[32;01m%s\033[0m\n"






