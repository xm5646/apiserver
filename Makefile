BINARY="api-server"
VERSION=1.0.0
BUILD=`date +%FT%T%z`
GIT_VERSION=$(shell pwd |git rev-parse --short HEAD)
IMAGE_TAG ?= "latest"
PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
PUB_SERVER=47.92.202.208
PUB_PORT=11167

default:
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/${BINARY} -tags=jsoniter

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	gofmt -s -w ${GOFILES}

fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

swag:
	swag init

install:
	@go mod tidy

run: swag fmt vet
	go run ./main.go

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	go vet $(VETPACKAGES)

docker-publish: docker-build docker-push
	@ssh -p ${PUB_PORT} root@${PUB_SERVER} "/root/update-dce.sh ${IMAGE_TAG}.tar"
	@say "发布完成"

image-upload: docker-build image-save
	@scp -P ${PUB_PORT} ${IMAGE_TAG}.tar  root@${PUB_SERVER}:/root/

image-save:
	@docker save -o  ${IMAGE_TAG}.tar 192.168.31.136/dashuo_containers/${BINARY}:${IMAGE_TAG}

linux-build: swag fmt vet
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/${BINARY} -tags=jsoniter

docker-build: linux-build
	@docker build --build-arg BINARY=${BINARY} -t 192.168.31.136/dashuo_containers/${BINARY}:${IMAGE_TAG} .

docker-push:
	@docker push 192.168.31.136/dashuo_containers/dce:${IMAGE_TAG}

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: default fmt fmt-check install test vet docker clean

