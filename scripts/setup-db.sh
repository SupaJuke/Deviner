#!/bin/bash

# Set default value for dbname
docker_name=postgres

if [ $# -gt 0 ]
then
	docker_name=$1
fi

# Copy schema.sql to container
docker cp scripts/schema.sql $docker_name:.

# Setup db
docker exec postgres chmod 777 schema.sql
docker exec -u postgres postgres psql -f schema.sql
