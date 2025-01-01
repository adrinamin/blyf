#!/bin/bash

podman build -t blyf:dev .

podman run --rm -d -p 8080:8080 blyf:dev # remove container after execution

# optional cleanup of dangling images
# podman image prune -f 
