#!/bin/bash

# ----------- Configuration -----------
API_URL="http://localhost:8080/api/login"  # تأكد من /api/ إذا مستعمل
EMAIL="mossab@example.com"
PASSWORD="strongpassword123"

# ----------- Login Request -----------
response=$(curl -s -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "'"$EMAIL"'",
    "password": "'"$PASSWORD"'"
  }')

echo "📨 Login response: $response"

# ----------- Extract token -----------
token=$(echo "$response" | jq -r '.token // .session_token // .access_token // ""')

# ----------- Check & Save token -----------
if [ -n "$token" ] && [ "$token" != "null" ]; then
  echo "$token" > session_token.txt
  echo "✅ Token enregistré dans session_token.txt"
else
  echo "❌ Erreur: Impossible d’extraire le token."
  echo "📦 Champs JSON disponibles:"
  echo "$response" | jq 'keys'
fi
