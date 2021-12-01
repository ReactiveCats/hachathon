CMD_PATH = ./cmd/server/main.go
GO = go
WIRE = wire
SWAG = swag
OUTPUT_FILENAME = main

default: help

help:
	@echo 'Usage:'
	@echo '		make build		build binary'
	@echo '		make install-deps	install dependencies (swag, ent, wire)'
	@echo '		make openapi-generate	generate openapi documentation'
	@echo '		make ent-generate	generate ent files'
	@echo '		make wire-generate	generate wire files'

build:
	@CGO_ENABLED=1 GOPROXY='https://proxy.golang.org,direct' ${GO} build -tags=nomsgpack -o ${OUTPUT_FILENAME} ${CMD_PATH}

openapi-generate:
	${SWAG} i --exclude internal/ent --parseVendor -g ${CMD_PATH}

ent-generate:
	@${GO} generate ./internal/ent/generate.go

wire-generate:
	@${WIRE} ./internal/bootstrap

install-deps:
	${GO} get entgo.io/ent/cmd/ent@latest
	${GO} get github.com/swaggo/swag/cmd/swag@latest
