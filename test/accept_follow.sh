#!/bin/bash

API_URL="http://localhost:8080/api/follow/accept"

#must session to be user 2
if [ ! -f session_token.txt ]; then
  echo "Error: session_token.txt not found"
  exit 1
fi

TOKEN=$(cat session_token.txt | tr -d '[:space:]')

FOLLOWER_ID=1
FOLLOWING_ID=2

response=$(curl -s -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -b "session_token=$TOKEN" \
  -d '{"follower_id": '"$FOLLOWER_ID"', "following_id": '"$FOLLOWING_ID"'}'
)

echo "Response: $response"
