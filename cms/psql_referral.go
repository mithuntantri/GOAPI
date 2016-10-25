package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func checkReferralID(referral_id string) bool {
  fmt.Println("Checking for valid referral_id:", referral_id)
  var count int8
  if referral_id == ""{
    return true
  }
  db.QueryRow("SELECT COUNT(*) FROM referral WHERE referral_id=$1",referral_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}

func updateReferralTable(referral_id string) string{
  var count int8
  db.QueryRow("SELECT referral_count FROM referral WHERE referral_id=$1",referral_id).Scan(&count)
  count++
  db.QueryRow("UPDATE referral SET referral_count=$1 where referral_id=$2", count, referral_id)
  var wallet_id string
  db.QueryRow("SELECT wallet_id FROM referral WHERE referral_id=$1",referral_id).Scan(&wallet_id)
  return wallet_id
}

func createReferralID(referral_id, wallet_id string)  bool{
  fmt.Println("Creating ReferralID",referral_id, wallet_id)
  db.QueryRow("INSERT INTO referral(referral_id, referral_count, wallet_id) VALUES($1,$2,$3);",referral_id, 0, wallet_id)
  return true
}
