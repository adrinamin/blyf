#!/bin/bash

docker build -t blyf:dev .

docker run --rm blyf:dev # remove container after execution
