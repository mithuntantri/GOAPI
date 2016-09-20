package main

import (
  "fmt"
  "database/sql"
  "github.com/metakeule/fmtdate"
  _ "github.com/lib/pq"
)
var db *sql.DB
func connectPSQL() {
    db, _ = sql.Open("postgres", "user=postgres dbname=mithun sslmode=disable password=postgres")
    defer db.Close()
}

func checkNewUser(mobileno string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
  if count == 0 {
    return true
  }else{
    return false
  }
}

func checkReferralID(referral_id string) bool {
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM referral WHERE referral_id=$1",referral_id).Scan(&count)
  if referral_id == "nil"{
    return true
  }
  if count == 0 {
    return false
  }else{
    return true
  }
}
func checkWalletID(wallet_id string) bool {
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM wallet WHERE wallet_id=$1",wallet_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func updateReferralTable(referral_id string) {
  var count int8
  db.QueryRow("SELECT referral_count FROM referral WHERE referral_id=$1",referral_id).Scan(&count)
  count++
  db.QueryRow("UPDATE referral SET referral_count=$1 where referral_id=$2", count, referral_id)
}

func addtoCredentials(mobileno, client_id, password string) bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO credentials(mobileno,client_id,password) VALUES($1,$2,$3);",
        mobileno, client_id, password).Scan(&lastInsertId)
  checkErr(err)
  fmt.Println("last inserted id =", lastInsertId)
  return true
}
func checkCredentials(mobileno, client_id, password string) bool{
  var hashedPass string
  var db_clientid string
  err := db.QueryRow("SELECT password, client_id FROM credentials WHERE mobileno=$1",
      mobileno).Scan(&hashedPass, &db_clientid)
  checkErr(err)
  if client_id == db_clientid && hashedPass == password{
    return true
  }else{
    return false
  }
}
func createReferralID(referral_id string)  bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO referral(referral_id, referral_count) VALUES($1,$2);",
        referral_id, 0).Scan(&lastInsertId)
        checkErr(err)
  fmt.Println("last inserted id =", lastInsertId)
  return true
}
func createWalletID(wallet_id string)  bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO wallet(wallet_id, referral_credits, profile_credits, promo_credits) VALUES($1,$2,$3,$4);",
        wallet_id, 0, 0 , 0).Scan(&lastInsertId)
        checkErr(err)
  fmt.Println("last inserted id =", lastInsertId)
  return true
}
func createProfile(request profileRequest, verified bool, referral_id, wallet_id, referred_id string)  bool{
  var lastInsertId string
  dob, _ := fmtdate.Parse("DD/MM/YYYY", request.Dob)
  err := db.QueryRow("INSERT INTO profile(mobileno, email_id, client_id, first_name, last_name, dob, gender, address, street, pin_code, verified, refferral_id, referred_id, wallet_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);",
        request.Mobileno, request.EmailID, request.ClientID, request.FirstName, request.LastName, dob, request.Gender, request.Address, request.Street, request.PinCode, verified, referral_id, referred_id, wallet_id).Scan(&lastInsertId)
  checkErr(err)
  fmt.Println("Last inserted id = ", lastInsertId)
  return true
}
