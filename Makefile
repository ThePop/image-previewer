ifeq ($(OS),Windows_NT)
    GOOS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOOS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        GOOS := darwin
    endif
endif

.PHONY: build
build:
	GOOS=${GOOS} CGO_ENABLED=0 GOARCH=amd64 go build -v -trimpath -o ./bin/image-previewer ./cmd/image-previewer/main.go

.PHONY: run
run:
	docker-compose up -d

.PHONY: test
test:
	go test ./internal/... -v -race -count=20

.PHONY: integration-test
integration-test:
	make docker-up-test && go test ./test/...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: docker-build
docker-build:
	docker-compose -f docker-compose.yml build

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yml up -d --remove-orphans

.PHONY: docker-down
docker-down:
	docker-compose -f docker-compose.yml down --remove-orphans

.PHONY: docker-build-test
docker-build-test:
	docker-compose -f docker-compose-test.yml build

.PHONY: docker-up-test
docker-up-test:
	docker-compose -f docker-compose-test.yml up -d --remove-orphans

.PHONY: docker-down-test
docker-down-test:
	docker-compose -f docker-compose-test.yml down --remove-orphans