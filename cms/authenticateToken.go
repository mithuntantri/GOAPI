package main

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
)
func authenticateToken(ID, client_id, tokenString string, mobile_device bool) (bool, bool) {
  vertoken, _ := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  // checkErr(err)
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    client_id := claims["client_id"].(string)
    valid := checkTokenExists(id, client_id, mobile_device)
    if valid {
      verifyToken(id, tokenString, mobile_device)
      return false, true
    }
    return false, false
  }
  deleteToken(ID, mobile_device)
  return true, false
}
func deleteauthToken(tokenString string, mobile_device bool) string{
  vertoken, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
       return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }
   return mySigningKey, nil
  })
  if(err != nil){
    return err.Error()
  }
  if claims, ok := vertoken.Claims.(jwt.MapClaims); ok && vertoken.Valid {
    id := claims["id"].(string)
    client_id := claims["client_id"].(string)
    valid := checkTokenExists(id, client_id, mobile_device)
    if valid {
      verifyToken(id, tokenString, mobile_device)
      deleteToken(id, mobile_device)
      return "success"
    }
  }
  return "failed"
}
