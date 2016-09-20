package main

import (
  "github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context)  {
  var json struct {
    Username string `json:"username"`
  }
  if c.Bind(&json) == nil{
    deleteToken(json.Username)
    c.JSON(200, gin.H{
      "status":"success",
    })
  }
}
