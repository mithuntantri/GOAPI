package main

import (
  "github.com/gin-gonic/gin"
)
func resendSignupHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `json:"mobileno"`
  }
  if c.Bind(&request) == nil {
    registered, blocked, verified := checkRegistrationExists(request.Mobileno)
    if registered && !blocked{
      blocked := callresOTP(request.Mobileno, "r")
      updateRegistrations(request.Mobileno, blocked, false)
      c.JSON(200, gin.H{
        "status" : "success",
        "blocked" :  blocked,
      })
    }else if registered && verified{
      c.JSON(200, gin.H{
        "status" : "success",
        "verified" : verified,
      })
    }
  }
}
