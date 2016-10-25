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
  db.QueryRow("INSERT INTO wallet(wallet_id, referral_credits, profile_credits, promo_credits) VALUES($1,$2,$3,$4);", wallet_id, 0, 0 , 0)
  return true
}

func getWallet(wallet_id string) (int, int, int){
  var referral_credits int
  var profile_credits int
  var promo_credits int
  db.QueryRow("SELECT referral_credits, profile_credits, promo_credits FROM wallet WHERE wallet_id=$1", wallet_id).Scan(&referral_credits, &profile_credits, &promo_credits)
  return referral_credits, profile_credits, promo_credits
}

func updateWallet(wallet_id, field_name string)  bool{
  var count int
  db.QueryRow(fmt.Sprintf("SELECT %s FROM wallet WHERE wallet_id=$1",QuoteIdentifier(field_name)),wallet_id).Scan(&count)
  count = count + 100
  db.QueryRow(fmt.Sprintf("UPDATE wallet SET %s=$1 where wallet_id=$2",QuoteIdentifier(field_name)), count, wallet_id)
  return true
}
