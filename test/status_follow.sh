#!/bin/bash

API_URL="http://localhost:8080/api/follow/status"
FOLLOWER_ID=5        
FOLLOWING_ID=2      

response=$(curl -s -G "$API_URL" \
  --data-urlencode "follower_id=$FOLLOWER_ID" \
  --data-urlencode "following_id=$FOLLOWING_ID")

echo "📨 Follow status response: $response"
