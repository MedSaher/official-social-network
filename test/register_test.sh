#!/bin/bash


curl -X POST http://localhost:8080/api/register \
  -F "nickname=beta_user" \
  -F "email=beta@example.com" \
  -F "password=strongpassword123" \
  -F "firstName=Beta" \
  -F "lastName=User" \
  -F "gender=male" \
  -F "dateOfBirth=1995-09-10" \
  -F "aboutMe=Développeur backend motivé." \
  -F "privacyStatus=private"
