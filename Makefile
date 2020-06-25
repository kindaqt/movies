ENV_FILE=.env
# APP
APP_NAME=movies
APP_VERSION=0.0.1
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
# API
API_DIR=./api
API_BINARY_NAME=api
API_BINARY_PATH=$(API_DIR)/out/$(API_BINARY_NAME)
API_PORT=8080

all: test build
build: 
	export GIN_MODE=release && $(GOBUILD) -o $(API_BINARY_PATH) .
test: 
	$(GOTEST) -v $(API_DIR)/...
clean: 
	$(GOCLEAN)
	rm -f $(API_BINARY_NAME)
run:
	$(GOBUILD) -o $(API_BINARY_PATH) .
	$(API_BINARY_PATH)
deps:
	# $(GOGET) github.com/markbates/goth
hot-reload:
	CompileDaemon -build="go build -o $(API_BINARY_PATH)" -command="$(API_BINARY_PATH)" -log-prefix=false -color=true -exclude-dir=".git"

########################
# Database
########################
migrate-up:
	@export MIGRATION_CMD=up \
	&& export MIGRATION_VERSION=1 \
	&& docker-compose up migration
migrate-down:
	@export MIGRATION_CMD=down \
	&& export MIGRATION_VERSION=1 \
	&& docker-compose up migration
migrate-version:
	@export MIGRATION_CMD=version \
	&& export MIGRATION_VERSION=1 \
	&& docker-compose up migration
migrate-force:
	@export MIGRATION_CMD=force \
	&& export MIGRATION_VERSION=1 \
	&& docker-compose up migration

########################
# Docker
########################
API_DOCKER_CONTAINER_NAME=$(APP_NAME)-$(API_BINARY_NAME)

ENV_FILE=.env

docker-up: 
	@source .env && docker-compose up
# docker-build:
# 	docker build -f $(API_DIR)/Dockerfile -t $(API_DOCKER_CONTAINER_NAME):$(APP_VERSION) .
# docker-run:
# 	docker run -d -p $(API_PORT):$(API_PORT) --name $(API_DOCKER_CONTAINER_NAME) -it $(API_DOCKER_CONTAINER_NAME):$(APP_VERSION)
# docker-logs:
# 	docker logs -f $(API_DOCKER_CONTAINER_NAME)