package main

import (
  "fmt"
  "strings"
  "database/sql"
  // "github.com/metakeule/fmtdate"
  _ "github.com/lib/pq"
)

var db *sql.DB

func connectPSQL() {
    db, _ = sql.Open("postgres", "user=postgres dbname=mithun sslmode=disable password=postgres")
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
func checkWalletID(wallet_id string) bool {
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM wallet WHERE wallet_id=$1",wallet_id).Scan(&count)
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

func addtoCredentials(mobileno , client_id , password string) bool{
  fmt.Println(mobileno, client_id, password)
  db.QueryRow("INSERT INTO credentials (mobileno, client_id, password) VALUES($1,$2,$3);", mobileno, client_id, password)
  return true
}
func checkCredentials(mobileno, client_id, password string) bool{
  var hashedPass string
  var db_clientid string
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1", mobileno).Scan(&count)
  if count == 0{
    return false
  }
  err := db.QueryRow("SELECT password, client_id FROM credentials WHERE mobileno=$1",
      mobileno).Scan(&hashedPass, &db_clientid)
  checkErr(err)
  fmt.Println(client_id, db_clientid)
  fmt.Println(hashedPass, password)
  if password == ""{
    return true
  }
  if client_id == db_clientid && hashedPass == password{
    return true
  }else{
    return false
  }
}
func verifyMobileno(mobileno, client_id string) (bool, bool){
  var db_clientid string
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
  if count == 0{
    return false, false
  }
  err := db.QueryRow("SELECT client_id FROM credentials WHERE mobileno=$1",
      mobileno).Scan(&db_clientid)
  checkErr(err)
  if client_id == db_clientid{
    return true, true
  }else{
    return false , true
  }
}
func createReferralID(referral_id, wallet_id string)  bool{
  fmt.Println("Creating ReferralID",referral_id, wallet_id)
  db.QueryRow("INSERT INTO referral(referral_id, referral_count, wallet_id) VALUES($1,$2,$3);",referral_id, 0, wallet_id)
  return true
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
func createProfile(request profileRequest, referral_id, wallet_id, referred_id string)  bool{
  db.QueryRow("INSERT INTO profile(mobileno, email_id, client_id, first_name, last_name, gender, address, street, pin_code, referral_id, referred_id, wallet_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);",
        request.Mobileno, request.EmailID, request.ClientID, request.FirstName, request.LastName, request.Gender, request.Address, request.Street, request.PinCode, referral_id, referred_id, wallet_id)
  return true
}
func getProfile(mobileno string) (profileRequest, string, string){
  var profile profileRequest
  var referral_id string
  var wallet_id string
  db.QueryRow("SELECT mobileno, email_id, client_id, first_name, last_name, gender, address, street, pin_code, referral_id, wallet_id FROM profile WHERE mobileno=$1",mobileno).Scan(&profile.Mobileno, &profile.EmailID, &profile.ClientID, &profile.FirstName, &profile.LastName, &profile.Gender, &profile.Address, &profile.Street, &profile.PinCode, &referral_id, &wallet_id)
  return profile, referral_id, wallet_id
}
func insertEmailMap(mobileno, email_id string) bool{
  fmt.Println("Creating Email Map", mobileno, email_id)
  db.QueryRow("INSERT INTO emailid_map(email_id, mobileno) VALUES($1, $2);", email_id, mobileno)
  return true
}
func checkIfEmailID(email_id string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM emailid_map WHERE email_id=$1",email_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func QuoteIdentifier(name string) string {
    end := strings.IndexRune(name, 0)
    if end > -1 {
        name = name[:end]
    }
    return `"` + strings.Replace(name, `"`, `""`, -1) + `"`
}
func getMobileNumber(id, table_name, field_name string) string{
  var mobileno string
  db.QueryRow(fmt.Sprintf("SELECT mobileno FROM %s WHERE %s=$1",
  QuoteIdentifier(table_name),QuoteIdentifier(field_name)),id).Scan(&mobileno)
  return mobileno
}
func checkIfUsername(username string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM username_map WHERE username=$1",username).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func checkIfFBID(fb_id string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM fbid_map WHERE fb_id=$1",fb_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func insertUsernameMap(mobileno, username string)  bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO username_map(email_id, mobileno) VALUES($1, $2);", username, mobileno).Scan(&lastInsertId)
  checkErr(err)
  return true
}
func insertFbIDMap(mobileno, fbid string)  bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO fbid_map(email_id, mobileno) VALUES($1, $2);", fbid, mobileno).Scan(&lastInsertId)
  checkErr(err)
  return true
}
func updateWallet(wallet_id, field_name string)  bool{
  var count int
  db.QueryRow(fmt.Sprintf("SELECT %s FROM wallet WHERE wallet_id=$1",QuoteIdentifier(field_name)),wallet_id).Scan(&count)
  count = count + 100
  db.QueryRow(fmt.Sprintf("UPDATE wallet SET %s=$1 where wallet_id=$2",QuoteIdentifier(field_name)), count, wallet_id)
  return true
}
