package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func createProfile(request profileRequest, referral_id, wallet_id, referred_id string)  bool{
  fmt.Println("Creating Profile : ", referral_id, wallet_id, referral_id)
  db.QueryRow("INSERT INTO profile(mobileno, email_id, client_id, first_name, last_name, gender, referral_id, referred_id, wallet_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",
        request.Mobileno, request.EmailID, request.ClientID, request.FirstName, request.LastName, request.Gender, referral_id, referred_id, wallet_id)
  return true
}

func getProfile(mobileno string) (profileRequest, string, string){
  var profile profileRequest
  var referral_id string
  var wallet_id string
  db.QueryRow("SELECT mobileno, email_id, client_id, first_name, last_name, gender, referral_id, wallet_id FROM profile WHERE mobileno=$1",mobileno).Scan(&profile.Mobileno, &profile.EmailID, &profile.ClientID, &profile.FirstName, &profile.LastName, &profile.Gender, &referral_id, &wallet_id)
  return profile, referral_id, wallet_id
}

func updateProfile(request profileRequest) bool{
  db.QueryRow("UPDATE profile SET first_name=$2, last_name=$3, gender=$4 WHERE mobileno=$1",request.Mobileno, request.FirstName, request.LastName, request.Gender)
  return true
}
