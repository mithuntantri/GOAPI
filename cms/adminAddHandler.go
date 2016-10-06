package main

import (
  "github.com/gin-gonic/gin"
)
func adminAddHandler(c *gin.Context)  {
  // if c.Bind(&request) == nil {
  //   tokenString := c.Request.Header.Get("X-Authorization-Token")
  //   device := c.Request.Header.Get("Device")
  //   mobile_device := false
  //   if device == "mobile"{
  //     mobile_device = true
  //   }
  //   if tokenString == "" {
  //     c.JSON(200, gin.H{
  //       "status" : "failed",
  //       "message" : "Invalid token",
  //       "data":map[string]interface{}{},
  //     })
  //     return
  //   }
  //   expired, authorized := authenticateToken(request.Mobileno, request.ClientID, tokenString, mobile_device)
  //   if expired || !authorized{
  //     c.JSON(401, gin.H{
  //       "status" : "failed",
  //       "message" : "Invalid Token",
  //     })
  //   }
  // }
}
