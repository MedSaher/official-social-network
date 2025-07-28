#!/bin/bash

API_URL="http://localhost:8080/api/posts/fetch_posts"
TOKEN=$(cat session_token.txt)

curl -s -X GET "$API_URL" \
  -b "session_token=$TOKEN" | jq
