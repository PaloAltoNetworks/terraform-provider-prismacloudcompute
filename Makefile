HOSTNAME = paloaltonetworks.com
NAMESPACE = prismacloud
NAME = prismacloudcompute
BINARY = terraform-provider-${NAME}

VERSION ?= 0.0.1
OS_ARCH ?= darwin_amd64

default: install

format:
	gofmt -l -w .

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
