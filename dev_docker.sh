#!/bin/bash

container_id=$(docker ps -l -q)

if [ -n  "$container_id" ]; then
    docker stop "$container_id"

    # optional cleanup of dangling images
    docker image prune -f
fi

docker build -t blyf:dev .

docker run --rm -d -p 8080:8080 blyf:dev # remove container after execution
