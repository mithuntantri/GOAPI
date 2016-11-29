package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func insertEmailMap(mobileno, email_id string) bool{
  fmt.Println("Creating Email Map", mobileno, email_id)
  db.QueryRow("INSERT INTO emailid_map(email_id, mobileno) VALUES($1, $2);", email_id, mobileno)
  return true
}

func checkIfEmailID(email_id string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM credentials WHERE email_id=$1",email_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
