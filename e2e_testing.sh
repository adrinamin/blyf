#!/bin/bash

echo "uploading file"
curl -X POST -F "file=@ test.txt" http://localhost:8080/upload

echo "downloading file"
curl -X GET http://localhost:8080/download/test.txt


