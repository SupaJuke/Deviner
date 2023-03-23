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

	# Load env vars from .env
	set -a; source .env; set +a

	# Copy schema.sql to container
	docker cp migrations/schema.sql $CONTAINER_NAME:.

	# Setup db
	docker exec $CONTAINER_NAME chmod 777 schema.sql
	docker exec -u postgres $CONTAINER_NAME dropdb --if-exists $DB_NAME
	docker exec -u postgres $CONTAINER_NAME createdb $DB_NAME
	docker exec -u postgres $CONTAINER_NAME psql -d $DB_NAME -f schema.sql -v DB_SCHEMA=$DB_SCHEMA
else
	echo Docker is not installed. Please install Docker before using this script.
fi
