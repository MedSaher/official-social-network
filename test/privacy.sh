#!/bin/bash

API_URL="http://localhost:8080/api/profile/privacy"

if [ ! -f session_token.txt ]; then
  echo "Error: session_token.txt not found"
  exit 1
fi

TOKEN=$(cat session_token.txt | tr -d '[:space:]')

curl -s -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -b "session_token=$TOKEN" \
  -d '{"privacyStatus":"private"}' | jq
