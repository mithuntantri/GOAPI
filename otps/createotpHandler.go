package main

import (
  "github.com/gin-gonic/gin"
)
func createotpHandler(c *gin.Context)  {
  var request struct {
    EmailID string `json:"email_id"`
    Mobileno    string `json:"mobileno"`
    RequestType string `json:"request_type"`
  }
  if c.Bind(&request) == nil {
    if request.RequestType == "n" {
      if exists := checkExists(request.Mobileno); !exists {
        otp := generate_otp(request.Mobileno)
        if req := createOTP(request.Mobileno, otp); req {
          go sendOtpThroughMail(request.EmailID, otp)
          go expireOTPchannel(request.Mobileno)
        }
        c.JSON(200, gin.H{
          "status" : "success",
          "is_exists" : exists,
          "valid_request" : true,
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "success",
          "is_exists" : true,
          "valid_request" : true,
        })
      }
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "is_exists" : false,
        "valid_request" : false,
      })
    }
  }
}
