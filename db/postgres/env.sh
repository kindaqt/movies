export POSTGRESQL_HOST=localhost
export POSTGRESQL_PORT=5432
export POSTGRESQL_USER=docker
export POSTGRESQL_PASSWORD=docker
export POSTGRESQL_DB_NAME=docker
export POSTGRESQL_SSL=disable
export POSTGRESQL_URL="postgres://$POSTGRESQL_USER:$POSTGRESQL_PASSWORD@$POSTGRESQL_HOST:$POSTGRESQL_PORT/$POSTGRESQL_DB_NAME?sslmode=$POSTGRESQL_SSL"