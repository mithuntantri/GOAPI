package main

import (
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
      if !checkIfEmailID(request.ID){
        if !checkIfUsername(request.ID){
          if !checkIfFBID(request.ID){
            c.JSON(200, gin.H{
              "status" : "success",
              "valid" : false,
            })
          }else{
            //Credentials exists with FB ID
            mobileno = getMobileNumber(request.ID, "fbid_map", "fb_id")
          }
        }else{
          //Credentials Exists with Username
          mobileno = getMobileNumber(request.ID, "username_map", "username")
        }
      }else{
        //Credentials exists with Emailid
        mobileno = getMobileNumber(request.ID, "emailid_map", "email_id")
      }
      if checkCredentials(request.ID, request.ClientID, hashedPass){
        logintoken = generateToken(mobileno, request.ClientID, request.Expiry)
      }
    }else{
      //Credentials Exists with mobileno
      logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
    }
    isExists := checkTokenExists(logintoken.ID)
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
