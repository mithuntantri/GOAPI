package main

import (
  "github.com/gin-gonic/gin"
)

func resendotpHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `json:"mobileno"`
    RequestType string `json:"request_type"`
  }
  if c.Bind(&request) == nil {
    validnum := false
    validreq := false
    blocked := false
    status := "failed"
    if request.RequestType == "r"{
      validreq = true
      if exists := checkExists(request.Mobileno); exists {
        validnum = true
        blocked = resendOTP(request.Mobileno)
        status = "success"
      }else{
        status = "failed"
      }
    }
    c.JSON(200, gin.H{
      "status" : status,
      "is_blocked" : blocked,
      "valid_number": validnum,
      "valid_request" : validreq,
    })
  }
}
