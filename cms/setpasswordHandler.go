package main

import (
  "github.com/gin-gonic/gin"
)
func setpaswordHandler(c *gin.Context)  {
  var request struct{
    EmailID string `json:"email_id"`
    Mobileno string `json:"mobileno"`
    ClientID string `json:"client_id"`
    Password string `json:"password"`
  }
  if c.Bind(&request) == nil {
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    registered, blocked, verified := checkRegistrationExists(request.Mobileno)
    if registered && verified {
      hashedPass := bcryptPassword(request.Password)
      if addtoCredentials(request.Mobileno, request.EmailID, true, false, request.ClientID, hashedPass){
        logintoken := generateToken(request.Mobileno, request.ClientID, true)
        inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token, mobile_device)
        c.JSON(200, gin.H{
          "status":"success",
          "message":"",
          "data":map[string]interface{}{
            "validUser": true,
            "password_set" :  inserr,
            "profile_set" : false,
            "first_time_login" :  true,
            "secret":logintoken.Token,
          },
        })
      }
    }else if registered && !verified{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not verified",
        "data":map[string]interface{}{},
      })
    }else if registered && blocked{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number blocked for unsuccesfull verify attempts",
        "data":map[string]interface{}{},
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not registered",
        "data":map[string]interface{}{},
      })
    }
  }
}
