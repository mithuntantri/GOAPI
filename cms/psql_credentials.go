package main

import (
  "fmt"
  _ "github.com/lib/pq"
)
func checkNewUser(mobileno string) bool{
  fmt.Println("Checking if New user :", mobileno)
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE mobileno=$1",mobileno).Scan(&count)
  if count == 0 {
    return true
  }else{
    return false
  }
}

func addtoCredentials(mobileno , client_id , password string) bool{
  fmt.Println("Adding to Credentials", mobileno, client_id, password)
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
  password_verified := verifyBcrypt(hashedPass, password)
  if password == ""{
    return true
  }
  if client_id == db_clientid && password_verified{
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
