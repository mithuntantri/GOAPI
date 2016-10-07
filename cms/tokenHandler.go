package main

import (
  "time"
  "log"
  "github.com/dgrijalva/jwt-go"
)

func generateToken(ID, ClientID string, expiry bool) loginTokens{
  token := jwt.New(jwt.SigningMethodHS256)
  if expiry {
    token.Claims =  jwt.MapClaims{
      "id"  : ID,
      "client_id" : ClientID,
      "exp"   : time.Now().Add(time.Minute * 10).Unix(),
    }
  }else{
    token.Claims =  jwt.MapClaims{
      "id"  : ID,
      "client_id" : ClientID,
      "exp"   : time.Now().Add(time.Minute * 100).Unix(),
    }
  }
  tokenString, err := token.SignedString(mySigningKey)
  if(err != nil){
    log.Fatal(err)
  }
  //Strore login tokens for all users
  logintoken := loginTokens{
    ID : ID,
    ClientID: ClientID,
    Token: tokenString,
  }
  return logintoken
}
