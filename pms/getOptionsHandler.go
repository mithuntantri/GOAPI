package main

import (
  "github.com/gin-gonic/gin"
)
func getOptionsHandler(c *gin.Context)  {
  var request struct{
    hash string `json:"hash"`
    Choice string `json:"choice"`
  }
  if c.Bind(&request) == nil{
    // options_count, options := getProductTypes(request.Choice)
    c.JSON(200, gin.H{
      "status" : "success",
    })
  }
}
