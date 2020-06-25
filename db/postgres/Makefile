
###############################
## TODO: move all this into the parent makefile



SHELL_CMD=bash
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
PSQL_DOCKER_NAME=postgres:12
PSQL_CONTAINER_NAME=postgres
PSQL_CMD=psql

start:
	$(DOCKER_COMPOSE_CMD) up
start-quiet:
	$(DOCKER_COMPOSE_CMD) up -d
down:
	$(DOCKER_COMPOSE_CMD) down
clean:
	# TODO: cleanup commands
test:
	# TODO: ping test

###########################
# MIGRATE
###########################
migrate-up:
	migrate -path ./migrations -database ${POSTGRESQL_URL} up
migrate-down:
	migrate -path ./migrations -database ${POSTGRESQL_URL} down
migrate-version:
	migrate -path ./migrations -database ${POSTGRESQL_URL} version

###########################
# DEBUG
###########################
logs:
	$(DOCKER_COMPOSE_CMD) logs -f
shell:
	# TODO - connection command
	$(DOCKER_COMPOSE_CMD) exec $(PSQL_CONTAINER_NAME) $(SHELL_CMD)
login:
	$(DOCKER_COMPOSE_CMD) exec $(PSQL_CONTAINER_NAME) $(PSQL_CMD)


