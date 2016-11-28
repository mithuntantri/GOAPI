package main

import (
  "github.com/gin-gonic/gin"
)

func facebookHandler(c *gin.Context)  {
  var request struct{
    FacebookID string `form:"facebook_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    exists := false
    if checkIfFBID(request.FacebookID){
      exists = true
    }
    c.JSON(200, gin.H{
      "status" : "success",
      "exists" : exists,
    })
  }
}
