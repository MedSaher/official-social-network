#!/bin/bash

API_URL="http://localhost:8080/api/profile/another"
USERNAME="Ualpha"  

if [ ! -f session_token.txt ]; then
  echo "❌ session_token.txt not found. سجل الدخول الأول!"
  exit 1
fi

TOKEN=$(cat session_token.txt | tr -d '[:space:]')

response=$(curl -s -X GET "$API_URL?username=$USERNAME" \
  -H "Content-Type: application/json" \
  -b "session_token=$TOKEN"
)

echo "📨 Response:"
echo "$response" | jq
