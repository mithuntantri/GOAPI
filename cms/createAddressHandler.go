package main

import (
  "strings"
  "github.com/gin-gonic/gin"
)

type createAddress struct{
  Mobileno string `json:"mobileno"`
  Address string `json:"address"`
  Street string `json:"street"`
  PinCode string `json:"pin_code"`
  IsDefault bool `json:"is_default"`
}
func createAddressHandler(c *gin.Context)  {
  var request createAddress
  if c.Bind(&request) == nil{
    //Generate an Address ID
    first := strings.SplitN(request.Mobileno,"", 5)
    part1 := strings.ToUpper(first[0] + first[1] + first[2] + first[3])
    part2, _ := Generate(`[a-Z]{6}`)
    addressID := part1 + part2
    if checkAddressID(addressID) {
      part2, _ =Generate(`[a-Z]{6}`)
      addressID = part1 + part2
    }
    created := addAddress(addressID, request)
    if created{
      c.JSON(200, gin.H{
        "status" : "success",
        "message" : "Address Added Successfully",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "error",
        "messaged": "Failed to add address",
      })
    }
  }
}
