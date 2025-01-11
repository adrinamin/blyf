#!/bin/bash

echo "uploading file"
curl -X POST -F "file=@ test.txt" http://localhost:8080/upload

curl -X GET http://localhost:8080/blyf

echo "downloading file"
curl -X GET http://localhost:8080/download/test.txt

curl -X DELETE http://localhost:8080/delete/test.txt
