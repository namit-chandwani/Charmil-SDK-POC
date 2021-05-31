.DEFAULT_GOAL := build
HOST_BIN_FILE = host
PLUGIN_BIN_FILE = plugin

build:
	@go build -o "${HOST_BIN_FILE}" ./cmd/"${HOST_BIN_FILE}"
	@go build -o "${PLUGIN_BIN_FILE}" ./cmd/"${PLUGIN_BIN_FILE}"
run_host:
	./"${HOST_BIN_FILE}"
run_plugin:
	./"${PLUGIN_BIN_FILE}"
clean:
	go clean
	rm -f ${HOST_BIN_FILE} ${PLUGIN_BIN_FILE}