PUBLIC_REGISTRY_HOST=docker.io
PUBLIC_REGISTRY_OWNER=eecrocush
PUBLIC_REGISTRY_APP_NAME=module-39

BUILD_CMD=docker build -t
PUSH_CMD=docker push

CI_COMMIT_REF_NAME=latest

all: deps build test

deps:
	@go mod download
	@echo "Dependencies installed successfully"

build:
	go build ./

test:
	go test -v ./...

lint:
	golangci-lint run ./...

image:
	${BUILD_CMD} ${PUBLIC_REGISTRY_HOST}/${PUBLIC_REGISTRY_OWNER}/${PUBLIC_REGISTRY_APP_NAME}:${CI_COMMIT_REF_NAME} ./
	${PUSH_CMD} ${PUBLIC_REGISTRY_HOST}/${PUBLIC_REGISTRY_OWNER}/${PUBLIC_REGISTRY_APP_NAME}:${CI_COMMIT_REF_NAME}
	@echo "New ${PUBLIC_REGISTRY_HOST}/${PUBLIC_REGISTRY_OWNER}/${PUBLIC_REGISTRY_APP_NAME} image ready! Version: ${CI_COMMIT_REF_NAME}"
