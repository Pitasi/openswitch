APP=openswitch
MODULE=github.com/Pitasi/openswitch/cmd/openswitch
VERSION=0.0.1

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@mkdir -p bin
	go build -ldflags '-X main.Version=v$(VERSION)' -o bin/$(APP) $(MODULE)


.PHONY: get
## get: runs go get
get:
	go get ./...

.PHONY: run
## run: runs go run main.go
run:
	go run ${MODULE}

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	go clean -x $(MODULE)

.PHONY: test
## test: runs go test with default values
test: get
	go test -v -count=1 -race ./...

.PHONY: setup
## setup: setup go modules
setup:
	go mod init \
		&& go mod tidy \
		&& go mod vendor

# helper rule for deployment
check-environment:
ifndef ENV
	$(error ENV not set, allowed values - `staging` or `production`)
endif

#.PHONY: docker-build
### docker-build: builds the stringifier docker image to registry
#docker-build: build
#	docker build -t ${APP}:${COMMIT_SHA} .
#
#.PHONY: docker-push
### docker-push: pushes the stringifier docker image to registry
#docker-push: check-environment docker-build
#	docker push ${REGISTRY}/${ENV}/${APP}:${COMMIT_SHA}

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: "
	@sed -n 's/^## //p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
