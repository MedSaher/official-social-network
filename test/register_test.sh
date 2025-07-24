#!/bin/bash

curl -s -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "mossab_user",
    "email": "mossab@example.com",
    "password": "strongpassword123",
    "firstName": "Mossab",
    "lastName": "Lahbib",
    "gender": "male",
    "dateOfBirth": "1993-08-15",
    "aboutMe": "Je suis passionné de code et de backend Go.",
    "privacyStatus": "public"
  }'
