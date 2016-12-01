package main

import (
  "github.com/gin-gonic/gin"
)
func setValueHandler(c *gin.Context)  {
  var request struct{
    Hash string `json:"hash"`
    Choice int `json:"choice"`
    Option int `json:"option"`
  }
  if c.Bind(&request) == nil{
    updateHashTable(request.Hash, request.Choice, request.Option)
    c.JSON(200, gin.H{
      "status" : "success",
      "message" : "Option Changed Successfully",
      "data":map[string]interface{}{
        "total_price" : "699.00",
      },
    })
  }
}
