#!/bin/bash

podman build -t blyf:dev .

podman run --rm blyf:dev # remove container after execution
