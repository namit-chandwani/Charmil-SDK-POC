.DEFAULT_GOAL := build
HOST_BIN_FILE=host
EXTENSION_BIN_FILE=extension

build:
	@go build -o "${HOST_BIN_FILE}" ./cmd/"${HOST_BIN_FILE}"
	@go build -o "${EXTENSION_BIN_FILE}" ./cmd/"${EXTENSION_BIN_FILE}"
run_host:
	./"${HOST_BIN_FILE}"
run_extension:
	./"${EXTENSION_BIN_FILE}"
clean:
	go clean
	rm -f ${HOST_BIN_FILE} ${EXTENSION_BIN_FILE}