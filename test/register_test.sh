#!/bin/bash

curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "mossab93",
    "username": "mossab_user",
    "email": "mossab@example.com",
    "password": "strongpassword123",
    "firstName": "mossab",
    "lastName": "mossab",
    "gender": "Male",
    "dateOfBirth": "1993-08-15",
    "aboutMe": "Je suis passionné de code et de backend Go."
  }'
