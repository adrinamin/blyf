#!/bin/bash

container_id=$(podman ps -l -q)

if [ -z  "$container_id" ]; then
    echo "No running containers found."
    exit 1
fi

podman stop "$container_id"

# optional cleanup of dangling images
podman image prune -f

podman build -t blyf:dev .

podman run --rm -d -p 8080:8080 blyf:dev # remove container after execution
