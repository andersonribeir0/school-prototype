#!/bin/bash

URL="http://localhost:8080/courses"

curl -X POST "$URL" \
     -H "Content-Type: application/json" \
     -d @createCourse.json

echo
