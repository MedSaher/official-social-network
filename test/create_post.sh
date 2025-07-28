#!/bin/bash

API_URL="http://localhost:8080/api/posts/create_post"
TOKEN=$(cat session_token.txt)

curl -s -X POST "$API_URL" \
  -b "session_token=$TOKEN" \
  -F "content=Hello, this is my first post!" \
  -F "privacy=public" | jq
