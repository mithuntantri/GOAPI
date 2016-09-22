package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
)
func authenticateToken(ID, client_id, tokenString string) (bool, bool) {
  fmt.Println(tokenString)
  vertoken, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  checkErr(err)
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    client_id := claims["client_id"].(string)
    valid := checkTokenExists(id, client_id)
    if valid {
      verifyToken(id, tokenString)
      return false, true
    }
    return false, false
  }
  deleteToken(ID)
  return true, false
}
func deleteauthToken(tokenString string) string{
  vertoken, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  checkErr(err)
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    client_id := claims["client_id"].(string)
    valid := checkTokenExists(id, client_id)
    if valid {
      verifyToken(id, tokenString)
      deleteToken(id)
      return "success"
    }
  }
  return "failed"
}
