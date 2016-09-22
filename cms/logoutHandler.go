package main

import (
  "github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context)  {
  var json struct {
    ID string `json:"mobileno"`
  }
  if c.Bind(&json) == nil{
    deleteToken(json.ID)
    c.JSON(200, gin.H{
      "status":"success",
    })
  }
}
