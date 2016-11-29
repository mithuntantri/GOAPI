package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
)

func resendEmailVerificationHandler(c *gin.Context){
  var request struct{
    Token string `form:"token" binding:"required"`
  }
  if c.Bind(&request) == nil{
    if mobileno, client_id, email_id := getDetailsfromToken(request.Token); mobileno != "" || client_id != "" || email_id != ""{
      sendEmailVerification(mobileno, client_id, email_id)
      c.JSON(200, gin.H{
        "status" : "success",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Failed to Send Verification Email! Please try again",
      })
    }
  }
}

func getDetailsfromToken(tokenString string) (string, string, string) {
  vertoken, _ := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  if claims, _ := vertoken.Claims.(jwt.MapClaims); !vertoken.Valid {
    id := claims["id"].(string)
    email_id := claims["email_id"].(string)
    client_id := claims["client_id"].(string)
    fmt.Println("Resending email for :", id, email_id, client_id)
    return id, client_id, email_id
  }
  return "","",""
}
