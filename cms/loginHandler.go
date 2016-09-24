package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context)  {
  var request struct {
    ID string `json:"id"`
    Password string `json:"password"`
    ClientID string `json:"client_id"`
    Device string `json:"device"`
    Expiry bool `json:"set_expiry"`
  }
  if c.Bind(&request) == nil{
    var mobileno string
    var logintoken loginTokens
    hashedPass := getHashedPassword(request.Password)
    if !checkCredentials(request.ID, request.ClientID, hashedPass){
      fmt.Println("User exists")
      if !checkIfEmailID(request.ID){
        fmt.Println("User with email id exists")
        if !checkIfUsername(request.ID){
          fmt.Println("User with username id exists")
          if !checkIfFBID(request.ID){
            c.JSON(200, gin.H{
              "status" : "success",
              "valid" : false,
            })
            return
          }else{
            //Credentials exists with FB ID
            mobileno = getMobileNumber(request.ID, "fbid_map", "fb_id")
          }
        }else{
          //Credentials Exists with Username
          mobileno = getMobileNumber(request.ID, "username_map", "username")
          fmt.Println(mobileno)
        }
      }else{
        //Credentials exists with Emailid
        mobileno = getMobileNumber(request.ID, "emailid_map", "email_id")
        fmt.Println(mobileno)
      }
      if checkCredentials(mobileno, request.ClientID, hashedPass){
        logintoken = generateToken(mobileno, request.ClientID, request.Expiry)
      }
    }else{
      //Credentials Exists with mobileno
      logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
    }
    isExists := checkTokenExists(logintoken.ID, logintoken.ClientID)
    if !isExists {
      inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token)
      if inserr {
        c.JSON(200, gin.H{
          "data":map[string]interface{}{
              "validUser": true,
              "secret":logintoken.Token,
          },
          "message":"",
          "status":"success",
        })
      }
    }else{
      c.JSON(200, gin.H{
        "data":map[string]interface{}{
            "validUser": true,
            "secret": "",
          },
          "message":"",
          "status":"failed",
        })
      }
  }
}
