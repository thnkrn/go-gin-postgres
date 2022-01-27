BINARY_NAME=go-gin-postgres

build:
	go build

run:
	./${BINARY_NAME}

serve:
	make build
	make run

watch:
	fresh