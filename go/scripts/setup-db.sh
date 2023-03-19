#!/bin/bash
docker --version > /dev/null 2>&1
if [[ $? -eq 0 ]]
then
	if [[ $# -eq 0 ]]
	then
		echo 'setup-db.sh <PostgreSQL container name>'
		exit 1
	else
		CONTAINER_NAME=$1
	fi

	# Set default value for db_name
	DB_NAME=pooe_game_db

	source ./.env

	# Copy schema.sql to container
	docker cp scripts/schema.sql $CONTAINER_NAME:.

	# Setup db
	docker exec $CONTAINER_NAME chmod 777 schema.sql
	docker exec -u postgres $CONTAINER_NAME dropdb --if-exists $DB_NAME
	docker exec -u postgres $CONTAINER_NAME createdb $DB_NAME
	docker exec -u postgres $CONTAINER_NAME psql -d $DB_NAME -f migrations/schema.sql
else
	echo Docker is not installed. Please install Docker before using this script.
fi