#!/bin/bash

container_id=$(podman ps -l -q)

if [ -n  "$container_id" ]; then
    podman stop "$container_id"

    # optional cleanup of dangling images
    podman image prune -f
fi

podman build -t blyf:dev .

podman run --rm -d -p 8080:8080 blyf:dev # remove container after execution
