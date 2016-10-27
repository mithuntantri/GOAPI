package main

import (
  "github.com/gin-gonic/gin"
)
func updateAddressHandler(c *gin.Context)  {
  var request address
  if c.Bind(&request) == nil{
    if checkAddressID(request.AddressID) {
      updated := updateAddress(request)
      if updated{
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Address Updated Successfully",
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "error",
          "message": "Failed to update address",
        })
      }
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Invalid Address ID",
      })
    }
  }
}
