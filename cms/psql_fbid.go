package main

import (
  "fmt"
  _ "github.com/lib/pq"
)

func checkIfFBID(fb_id string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM fbid_map WHERE fb_id=$1",fb_id).Scan(&count)
  fmt.Println("Checking Facebook ID Count:",count,fb_id)
  if count == 0 {
    return false
  }else{
    return true
  }
}

func insertFbIDMap(mobileno, fbid string)  bool{
  db.QueryRow("INSERT INTO fbid_map(fb_id, mobileno) VALUES($1, $2);", fbid, mobileno)
  return true
}
