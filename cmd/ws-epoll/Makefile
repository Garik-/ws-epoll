include ../../includes.mk

APP := ws-epoll
ID := 1

.PHONY: run build

build:
	@echo "=> building $(APP) binary"
	$(GO) build -trimpath -ldflags "${GO_LDFLAGS}" -o $(BIN_DIR)/$(APP) .

run: build
	@echo "=> run $(APP) $(VERSION)"
	@$(BIN_DIR)/$(APP)