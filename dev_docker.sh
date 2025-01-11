#!/bin/bash

container_id=$(docker ps -l -q)

if [ -z  "$container_id" ]; then
    echo "No running containers found."
    exit 1
fi

docker stop "$container_id"

# optional cleanup of dangling images
docker image prune -f

docker build -t blyf:dev .

docker run --rm -d -p 8080:8080 blyf:dev # remove container after execution

# optional cleanup of dangling images
# docker image prune -f 
