package main

import (
  _ "github.com/lib/pq"
)

func checkIfUsername(username string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM username_map WHERE username=$1",username).Scan(&count)
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
