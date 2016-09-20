package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
)

func authenticateToken(ID, tokenString string) (bool, bool) {
  token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  checkErr(err)
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    id := claims["ID"].(string)
    valid := checkTokenExists(id)
    if valid {
      verified := verifyToken(id, tokenString)
      return false, verified
    }
  }
  deleteToken(ID)
  return true, false
}
