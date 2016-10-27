package main

import (
  "github.com/gin-gonic/gin"
)

func deleteAddressHandler(c *gin.Context)  {
  var request struct{
    AddressID string `form:"address_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    if checkAddressID(request.AddressID) {
      deleted := deleteAddress(request.AddressID)
      if deleted{
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Address deleted Successfully",
        })
        }else{
          c.JSON(200, gin.H{
            "status" : "error",
            "message": "Failed to delete address",
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
