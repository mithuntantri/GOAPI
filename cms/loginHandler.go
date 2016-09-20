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
    if len(request.ID) == 10{
      hashedPass := getHashedPassword(request.Password)
      if valid := checkCredentials(request.ID, request.ClientID, hashedPass); valid{
        var logintoken loginTokens
        logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
        isExists := checkTokenExists(logintoken.ID)
        fmt.Println("isExists:", isExists)
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
    }
  }
