#!/bin/bash

API_URL="http://localhost:8080/api/profile"

if [ ! -f session_token.txt ]; then
  echo "Error: session_token.txt not found"
  exit 1
fi

TOKEN=$(cat session_token.txt | tr -d '[:space:]')

USER_ID=5

response=$(curl -s -X GET "$API_URL?user_id=$USER_ID" \
  -H "Content-Type: application/json" \
  -b "session_token=$TOKEN"
)

echo "📨 Response:"
echo "$response" | jq
