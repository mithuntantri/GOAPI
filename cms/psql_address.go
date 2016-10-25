package main

import (
  _ "github.com/lib/pq"
)

func checkAddressID(address_id string)  bool{
  var count int8
  db.QueryRow("SELECT COUNT(*) FROM address WHERE address_id=$1",address_id).Scan(&count)
  if count == 0 {
    return false
  }else{
    return true
  }
}
func updateAddress(address_id string, request createAddress) bool{
  if request.IsDefault{
    err := db.QueryRow("UPDATE address SET is_default=$1 WHERE mobileno=$2;",false,request.Mobileno)
    if err != nil{
      return false
    }
  }
  err := db.QueryRow("UPDATE address SET address=$1, street=$2, pin_code=$3, is_default=$4 WHERE address_id=$5;", request.Address, request.Street, request.PinCode, request.IsDefault, address_id)
  if err != nil{
    return false
  }
  return true
}
func addAddress(address_id string, request createAddress) bool{
  err := db.QueryRow("INSERT INTO address (address_id, mobileno, address, street, pin_code, is_default) VALUES ($1, $2, $3, $4, $5, $6);", address_id, request.Mobileno, request.Address, request.Street, request.PinCode, request.IsDefault)
  if err != nil{
    return false
  }
  return true
}
func fetchAddresses(mobileno string)  Addresses{
  var all_addresses Addresses
  all_addresses.AllAddresses = make([]address,0)
  rows,err := db.Query("SELECT * from address WHERE mobileno=$1", mobileno)
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var row address
    if err := rows.Scan(&row.AddressID, &row.Mobileno, &row.Address, &row.Street, &row.PinCode, &row.IsDefault); err != nil{
      checkErr(err)
    }else{
      all_addresses.AllAddresses = append(all_addresses.AllAddresses, row)
    }
  }
  return all_addresses
}
