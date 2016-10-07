package main

import (
  "github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context)  {
  tokenString := c.Request.Header.Get("X-Authorization-Token")
  device := c.Request.Header.Get("Device")
  mobile_device := false
  if device == "mobile"{
    mobile_device = true
  }
  if tokenString == "" {
    c.JSON(401, gin.H{
      "status" : "failed",
      "message" : "Request Unauthorized",
    })
    return
  }
    result := deleteauthToken(tokenString, mobile_device)
    var message string
    if result == "success"{
      message = "Logged Out Successfully"
    }else{
      message = "Request failed"
    }
    c.JSON(200, gin.H{
      "status" : result,
      "message" : message,
    })
}
