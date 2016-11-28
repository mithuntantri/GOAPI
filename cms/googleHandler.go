package main

import (
  "github.com/gin-gonic/gin"
)

func googleHandler(c *gin.Context)  {
  var request struct{
    GoogleID string `form:"google_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    exists := false
    if checkIfEmailID(request.GoogleID){
      exists = true
    }
    c.JSON(200, gin.H{
      "status" : "success",
      "exists" : exists,
    })
  }
}
