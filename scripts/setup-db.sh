#!/bin/bash
if [[ $(docker --version) ]]
then
	echo uwu
fi

# Set default value for container name
docker_name=postgres
db_name=pooe_game_db

if [ $# -gt 0 ]
then
	docker_name=$1
fi

# Copy schema.sql to container
docker cp scripts/schema.sql $docker_name:.

# Setup db
docker exec $docker_name chmod 777 schema.sql
docker exec -u postgres $docker_name dropdb --if-exists $db_name
docker exec -u postgres $docker_name createdb $db_name
docker exec -u postgres $docker_name psql -d $db_name -f schema.sql
