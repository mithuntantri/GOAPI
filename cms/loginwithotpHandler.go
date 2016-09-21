package main

import (
  "github.com/gin-gonic/gin"
)

func loginwithotpHanlder(c *gin.Context)  {
  var request struct{
    Mobileno string `json:"mobileno"`
    ClientID string `json:"cleint_id"`
    RequestType string `json:"request_type"`
    Otp string `json:"otp"`
  }
  if c.Bind(&request) == nil {
    validm, validc := verifyMobileno(request.Mobileno, request.ClientID)
    if validm && validc {
      if request.RequestType == "n"{
        response := callnewOTP(request.Mobileno, "n")
        c.JSON(200, gin.H{
          "status": "success",
          "valid_mobileno" : validm,
          "valid_client_id": validc,
          "otp_generated": response,
        })
      }else if request.RequestType == "r"{
        blocked := callresOTP(request.Mobileno, "r")
        c.JSON(200, gin.H{
          "status" : "success",
          "valid_mobileno" : validm,
          "valid_client_id": validc,
          "blocked" :  blocked,
        })
      }else if request.RequestType == "v"{
        blocked, verified := callverOTP(request.Mobileno, request.Otp, "v")
        c.JSON(200, gin.H{
          "status" : "success",
          "blocked" : blocked,
          "verified" : verified,
          "valid_mobileno" : validm,
          "valid_client_id": validc,
        })
      }
    }
    c.JSON(200, gin.H{
      "status" : "success",
      "valid_mobileno" : validm,
      "valid_client_id" : validc,
    })
  }
}
