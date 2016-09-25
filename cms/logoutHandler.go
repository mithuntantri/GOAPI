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
  result := deleteauthToken(tokenString, mobile_device)
  c.JSON(200, gin.H{
    "status" : result,
    "message" : "",
    "data":map[string]interface{}{},
  })
}
