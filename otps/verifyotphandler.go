package main

import (
  "github.com/gin-gonic/gin"
)

func verifyotpHandler(c *gin.Context)  {
  var request struct {
    Mobileno    string `json:"mobileno"`
    Otp         string `json:"otp"`
    RequestType string `json:"request_type"`
  }
  if c.Bind(&request) == nil {
    validnum := false
    validreq := false
    verified := false
    blocked := false
    status := "failed"
    if request.RequestType == "v"{
      validreq = true
      if exists := checkExists(request.Mobileno); exists {
        validnum = true
        verified, blocked = verifyOTP(request.Mobileno, request.Otp)
        status = "success"
      }else{
        status = "failed"
      }
    }
    c.JSON(200, gin.H{
      "status" : status,
      "verified" : verified,
      "is_blocked" : blocked,
      "valid_number": validnum,
      "valid_request" : validreq,
    })
  }
}
