#!/bin/bash

curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
  "nickname": "testnick",
  "username": "testuser",
  "email": "test@example.com",
  "password": "secret123",
  "firstName": "John",
  "lastName": "Doe",
  "gender": "male",
  "dateOfBirth": "1990-01-01",
  "aboutMe": "Just testing.",
  "privacyStatus": "public"
}
'
