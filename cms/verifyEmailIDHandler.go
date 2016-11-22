package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
)

func verifyEmailIDHandler(c *gin.Context)  {
  var request struct{
    Token string `form:"token" binding:"required"`
  }
  if c.Bind(&request) == nil{
    if request.Token == ""{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Invalid Request! Failed to Verify Email ID",
      })
    }else if mobileno, verified := authenticateEmailToken(request.Token); verified{
      if updateVerifiedEmail(mobileno){
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Email ID Verified Successfully",
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Something went wrong! Failed to verify Email ID",
        })
      }
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Failed to verify Email ID",
      })
    }
  }
}
func authenticateEmailToken(tokenString string) (string, bool) {
  vertoken, _ := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    // email_id := claims["email_id"].(string)
    client_id := claims["client_id"].(string)
    valid := checkTokenExists(id, client_id, true)
    if valid {
      if verifyEmailToken(id, tokenString){
        return id, true
      }
    }
    return "", false
  }
  return "", false
}
