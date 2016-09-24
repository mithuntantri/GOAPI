package main

import (
  "github.com/gin-gonic/gin"
)

func verifySignupHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `json:"mobileno"`
    Otp string `json:"otp"`
  }
  registered, blocked, verified := false, false, false
  if c.Bind(&request) == nil {
    //Check for the new Registrations
    registered, blocked, verified = checkRegistrationExists(request.Mobileno)
    if registered && !blocked && !verified {
      //Verify OTP
      blocked, verified = callverOTP(request.Mobileno, request.Otp, "v")
        //Update Registrations
        updateRegistrations(request.Mobileno, blocked, verified)
        c.JSON(200, gin.H{
          "status" : "success",
          "blocked" : blocked,
          "password_set" : false,
          "profile_set" : false,
          "first_time_login" : true,
          "verified" : verified,
        })
      }
    }else if !registered {
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not requested for OTP",
      })
    }else if registered && blocked {
      c.JSON(200, gin.H{
        "status": "failed",
        "message": "Mobile Number blocked beacuse of unsuccessful verify attempts",
      })
    }else if registered && verified{
      c.JSON(200, gin.H{
        "status": "failed",
        "message": "Mobile Number already registered & verified",
      })
    }
  }
