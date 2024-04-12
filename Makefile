VERSION=$(shell git describe --tags --abbrev=0)

ifeq ($(VERSION),)
	VERSION="0.0.0"
endif

lint:
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run --fix

test:
	go test -cover -race -v ./...

.PHONY: build
build:
	mkdir -p ./build
	go build \
		-o ./build/payment-processor ./cmd

image:
	docker build \
    		--tag rss3-network/payment-processor:$(VERSION) \
    		.

run:
	ENVIRONMENT=development go run ./cmd

.PHONY: genoapi
OAPI_SPEC ?= docs/gateway.yml
OAPI_TARGET ?= internal/service/hub/gen/oapi/
OAPI_TARGET_FILENAME ?= oapi.go
genoapi:
	mkdir -p $(OAPI_TARGET)
	go get github.com/deepmap/oapi-codegen/v2
	go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -package oapi -generate=types,client,server,spec,skip-prune -o $(OAPI_TARGET)$(OAPI_TARGET_FILENAME) $(OAPI_SPEC)
	go mod tidy

.PHONY: genmigration applymigration
MIG ?= new_migration
ENV ?= dev
genmigration:
	docker compose -f ./deploy/docker-compose.migration.yml up -d
	atlas migrate diff $(MIG) --env $(ENV)
	docker compose -f ./deploy/docker-compose.migration.yml down -v
applymigration:
	#atlas migrate apply --env $(ENV)
	echo "Please use goose instead"
