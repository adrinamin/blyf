#!/bin/bash

echo "TEST 1: happy path"
echo "STEP 1: upload file"
curl -X POST -F "file=@ test.pdf" http://localhost:8080/upload

curl -X GET http://localhost:8080/blyf

echo "STEP 2: download file"
curl -X GET http://localhost:8080/download/test.pdf

echo "STEP 3: delete file"
curl -X DELETE http://localhost:8080/delete/test.pdf


echo "TEST 2: wrong file extension"
curl -X POST -F "file=@ test.txt" http://localhost:8080/upload
