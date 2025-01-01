#!/bin/bash

docker build -t blyf:dev .

docker run --rm -d -p 8080:8080 blyf:dev # remove container after execution

# optional cleanup of dangling images
# docker image prune -f 
