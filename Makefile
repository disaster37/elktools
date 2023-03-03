TEST?=./...
KIBANA_URL ?= http://127.0.0.1:5601
ELASTICSEARCH_URL ?= http://127.0.0.1:9200
ELASTICSEARCH_USERNAME ?= elastic
ELASTICSEARCH_PASSWORD ?= changeme

all: help

build: fmt
	go build .

test: fmt
	KIBANA_URL=${KIBANA_URL} ELASTICSEARCH_URL=${ELASTICSEARCH_URL} ELASTICSEARCH_USERNAME=${ELASTICSEARCH_USERNAME} ELASTICSEARCH_PASSWORD=${ELASTICSEARCH_PASSWORD} go test $(TEST) -v -count 1 -parallel 1 -race -coverprofile=coverage.txt -covermode=atomic $(TESTARGS) -timeout 120m

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./
