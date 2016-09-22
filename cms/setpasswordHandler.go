package main

import (
  // "fmt"
  // "reflect"
  "github.com/gin-gonic/gin"
)
func setpaswordHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `json:"mobileno"`
    ClientID string `json:"client_id"`
    Password string `json:"password"`
  }
  if c.Bind(&request) == nil {
    registered, blocked, verified := checkRegistrationExists(request.Mobileno)
    if registered && verified {
      hashedPass := getHashedPassword(request.Password)
      // for i := 0; i < len(hashedPass); i++ {
      //   pass := string(hashedPass[i])
      //   fmt.Println(pass)
      //   fmt.Printf("%x ", hashedPass[i])
      // }
      if addtoCredentials(request.Mobileno, request.ClientID, hashedPass){
        logintoken := generateToken(request.Mobileno, request.ClientID, true)
        inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token)
        c.JSON(200, gin.H{
          "data":map[string]interface{}{
            "validUser": true,
            "password_set" :  inserr,
            "profile_set" : false,
            "first_time_login" :  true,
            "secret":logintoken.Token,
          },
          "message":"",
          "status":"success",
        })
      }
    }else if registered && !verified{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not verified",
      })
    }else if registered && blocked{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number blocked for unsuccesfull verify attempts",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not registered",
      })
    }
  }
}
