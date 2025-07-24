#!/bin/bash

API_URL="http://localhost:8080/api/follow/following"
USER_ID=2

curl -s -G "$API_URL" \
  --data-urlencode "user_id=$USER_ID" | jq .
