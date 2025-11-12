BINARY_NAME=rkt.exe

# Detect OS
ifeq ($(OS),Windows_NT)
    BINARY_PATH=.\tmp\$(BINARY_NAME)
    RM=del
    KILL=taskkill /IM $(BINARY_NAME) /F
    RUN_CMD=.\tmp\$(BINARY_NAME)
else
    BINARY_PATH=./tmp/$(BINARY_NAME)
    RM=rm -f
    KILL=pkill -f $(BINARY_NAME) || true
    RUN_CMD=./tmp/$(BINARY_NAME)
endif

## build: builds all binaries
build:
	@go mod vendor
	@go build -o $(BINARY_PATH) .
	@echo RKT built!

run: build
	@echo Starting RKT...
	@$(RUN_CMD) &
	@echo RKT started!

clean:
	@echo Cleaning...
	@go clean
	@$(RM) $(BINARY_PATH)
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run

stop:
	@echo "Stopping RKT..."
	@$(KILL)
	@echo Stopped RKT

restart: stop start