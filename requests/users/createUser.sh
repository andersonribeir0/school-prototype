#!/bin/bash

URL="http://localhost:8080/user"

curl -X POST "$URL" \
     -H "Content-Type: application/json" \
     -d @createUser.json

echo
