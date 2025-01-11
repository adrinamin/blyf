#!/bin/bash

# SELinux Contexts
# On systems with SELinux enabled, Podman requires proper labeling of mounted volumes to grant container processes the necessary access. 
# Without appropriate labels, SELinux may block access, resulting in permission errors.
# For that you need to include the :Z option when mounting the volume.
# This option tells Podman to label the content with a private unshared label, ensuring only the current container can use it.
podman run --rm -v "$(pwd)":/app:Z -w /app docker.io/library/golang:1-alpine "$@"

# --rm: removes the container after execution
# -v: mounts the volume
# -w sets the working directory
