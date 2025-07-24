#!/bin/bash

curl -s -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "beta_user",
    "email": "beta@example.com",
    "password": "strongpassword123",
    "firstName": "Beta",
    "lastName": "User",
    "gender": "male",
    "dateOfBirth": "1995-09-10",
    "aboutMe": "Développeur backend motivé.",
    "privacyStatus": "private"
  }'
