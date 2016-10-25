package main

import (
  _ "github.com/lib/pq"
)

func checkIfFBID(fb_id string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM fbid_map WHERE fb_id=$1",fb_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}

func insertFbIDMap(mobileno, fbid string)  bool{
  var lastInsertId string
  err := db.QueryRow("INSERT INTO fbid_map(email_id, mobileno) VALUES($1, $2);", fbid, mobileno).Scan(&lastInsertId)
  checkErr(err)
  return true
}
