#!/bin/bash

# Check if token file exists
if [ ! -f session_token.txt ]; then
  echo "❌ session_token.txt not found."
  exit 1
fi

# Read token and clean it
TOKEN=$(tr -d '[:space:]' < session_token.txt)

if [ -z "$TOKEN" ]; then
  echo "❌ Token is empty!"
  exit 1
fi

echo "📤 Sending logout request with token: $TOKEN"

# Send logout request
response=$(curl -s -X POST http://localhost:8080/api/logout \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json")

echo "📨 Logout response: $response"

# Check if logout was successful
success=$(echo "$response" | jq -r '.success // empty')
if [ "$success" == "true" ]; then
  echo "✅ Logout successful. Removing token."
  rm session_token.txt
else
  echo "❌ Logout failed. Details:"
  echo "$response" | jq
fi
