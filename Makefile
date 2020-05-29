# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=movies
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
	$(GOBUILD) -o ./out/$(BINARY_NAME) .
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o ./out/$(BINARY_NAME) .
	./out/$(BINARY_NAME)
deps:
	# $(GOGET) github.com/markbates/goth
hot-reload:
	CompileDaemon -build="go build -o ./out/$(BINARY_NAME)" -command="./out/$(BINARY_NAME)" -log-prefix=false -color=true -exclude-dir=".git"

DOCKER_CONTAINER_NAME=movies-api

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	# docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
	docker build -t $(DOCKER_CONTAINER_NAME):latest .
docker-run:
	docker run -d -p 8080:8080 --name $(DOCKER_CONTAINER_NAME) -it $(DOCKER_CONTAINER_NAME):latest
