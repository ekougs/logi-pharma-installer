SHELL := /bin/bash
BUILD_DIR="$(GOPATH)"/bin/logi-pharma-installer
VENDOR_DIR="$(BUILD_DIR)"/vendor
INSTALLER_BASE_NAME=logi-pharma-installer
ifeq ($(GOOS), windows)
    INSTALLER_NAME="$(INSTALLER_BASE_NAME)".exe
else
    INSTALLER_NAME="$(INSTALLER_BASE_NAME)"
endif

all: pkg_logi_pharma_installer clean

pkg_logi_pharma_installer:logi-pharma-installer pkg_resources
	@echo "Packaging installer..."
	@mkdir -p $(BUILD_DIR)
	@cp $(INSTALLER_NAME) $(BUILD_DIR)

pkg_resources:
	@echo "Packaging vendor resources..."
	@mkdir -p $(VENDOR_DIR)
	@cp -r vendor/ $(VENDOR_DIR)

logi-pharma-installer: logi-pharma-installer.go
	@echo "Building exec..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build

clean:
	@rm $(INSTALLER_NAME)
