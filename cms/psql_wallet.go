package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func checkWalletID(wallet_id string) bool {
  fmt.Println("Checking for valid wallet id:", wallet_id)
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM wallet WHERE wallet_id=$1",wallet_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}

func createWalletID(wallet_id string)  bool{
  db.QueryRow("INSERT INTO wallet(wallet_id, referral_credits, profile_credits, promo_credits, total_credits) VALUES($1,$2,$3,$4,$5);", wallet_id, 0, 0 ,0, 0)
  return true
}

func getWallet(wallet_id string) (float64, float64, float64, float64){
  var referral_credits float64
  var profile_credits float64
  var promo_credits float64
  var total_credits float64
  db.QueryRow("SELECT referral_credits, profile_credits, promo_credits, total_credits FROM wallet WHERE wallet_id=$1", wallet_id).Scan(&referral_credits, &profile_credits, &promo_credits, &total_credits)
  return referral_credits, profile_credits, promo_credits, total_credits
}

func updateWallet(wallet_id, field_name string)  bool{
  var count int
  db.QueryRow(fmt.Sprintf("SELECT %s FROM wallet WHERE wallet_id=$1",QuoteIdentifier(field_name)),wallet_id).Scan(&count)
  count = count + 100
  db.QueryRow(fmt.Sprintf("UPDATE wallet SET %s=$1 where wallet_id=$2",QuoteIdentifier(field_name)), count, wallet_id)
  return true
}
func useWallet(wallet_id string, used_amount float64) bool{
  var referral_credits float64
  var profile_credits float64
  var promo_credits float64
  var total_credits float64
  db.QueryRow("SELECT referral_credits, profile_credits, promo_credits, total_credits FROM wallet WHERE wallet_id=$1", wallet_id).Scan(&referral_credits, &profile_credits, &promo_credits, &total_credits)
  if profile_credits > 0 && profile_credits >= used_amount{
    profile_credits = profile_credits - used_amount
  }else if referral_credits > 0 && referral_credits >= used_amount{
    referral_credits = referral_credits - used_amount
  }else{
    promo_credits = promo_credits - used_amount
  }
  total_credits = total_credits - used_amount
  db.QueryRow("UPDATE wallet SET referral_credits=$1, profile_credits=$2, promo_credits=$3, total_credits=$4 WHERE wallet_id=$5", referral_credits, profile_credits, promo_credits, total_credits, wallet_id)
  return true
}
