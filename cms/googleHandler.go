package main

import (
  "github.com/gin-gonic/gin"
)

func googleHandler(c *gin.Context)  {
  var request struct{
    ClientID string `json:"clientId"`
    Code string `json:"code"`
    RedirectUri string `json:"redirectUri"`
    State string `json:"state"`
  }
  if c.Bind(&request) == nil{
    c.JSON(200, gin.H{
      "status" : "success",
    })
  }
}
