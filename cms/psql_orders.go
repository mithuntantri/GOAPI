package main

import (
  _ "github.com/lib/pq"
)

func checkOrderID(order_id string) bool{
  var count int8
  if order_id == ""{
    return true
  }
  db.QueryRow("SELECT COUNT(*) FROM orders WHERE order_id=$1",order_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func createOrder(order order) bool{
  return true
}
func isFirstOrder(mobileno string) bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM orders WHERE mobileno=$1", mobileno).Scan(&count)
  if(count == 0){
    return true
  }
  return false
}
