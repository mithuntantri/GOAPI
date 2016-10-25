package main

import (
  "github.com/gin-gonic/gin"
)
type address struct{
  AddressID string `json:"string"`
  Mobileno string `json:mobileno`
  Address string `json:"string"`
  Street string `json:"string"`
  PinCode string `json:"pin_code"`
  IsDefault bool `json:"is_default"`
}
type Addresses struct{
  AllAddresses []address `json:"all_addresses"`
}
func getAddressesHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno" binding:"required"`
  }
  if c.Bind(&request) == nil{
    var all_addresses Addresses
    all_addresses = fetchAddresses(request.Mobileno)
    c.JSON(200, gin.H{
      "status" :"success",
      "data": all_addresses,
    })
  }
}
