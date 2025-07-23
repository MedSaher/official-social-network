#!/bin/bash

# Test script for accepting a follow request

# Helper function to send POST requests
send_post_request() {
    local url=$1
    local data=$2
    curl -X POST -H "Content-Type: application/json" -d "$data" "$url"
}

# Test case: Accept a follow request successfully
echo "Test Case 1: Accept Follow Request"
response=$(send_post_request "http://localhost:8080/follow/accept" '{"follower_id": 1, "following_id": 2}')
echo "Response: $response"