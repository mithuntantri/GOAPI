package main

import (
  "strings"
  "github.com/gin-gonic/gin"
)

type profileRequest struct{
  Mobileno  string `json:"mobileno"`
  EmailID   string `json:"email_id"`
  ClientID  string `json:"client_id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Dob       string `json:"dob"`
  Gender    string `json:"gender"`
  Address   string `json:"address"`
  Street    string `json:"street"`
  PinCode   string `json:"pin_code"`
}
func setprofileHandler(c *gin.Context)  {
  var request profileRequest
  if c.Bind(&request) == nil {
    tokenString := c.Request.Header.Get("X-Authorization-Token")
    expired, authorized := authenticateToken(request.Mobileno, tokenString)
    registered, blocked, verified := checkRegistrationExists(request.Mobileno)
    if !expired && authorized && registered && verified{
      //Get Refferred ID
      referredID := getReferredID(request.Mobileno)

      //Generate a Referral ID
      first := strings.SplitN(request.FirstName,"", 5)
      part1 := strings.ToUpper(first[0] + first[1] + first[2] + first[3])
      part2, _ := Generate(`[a-Z]{6}`)
      referralID := part1 + part2
      if checkReferralID(referralID) {
        part2, _ =Generate(`[a-Z]{6}`)
        referralID = part1 + part2
      }
      createReferralID(referralID)

      //Generate a Wallet ID
      first = strings.SplitN(request.Mobileno,"", 5)
      part1 = strings.ToUpper(first[0] + first[1] + first[2] + first[3])
      part2, _ = Generate(`[a-Z]{6}`)
      walletID := part1 + part2
      if checkWalletID(walletID) {
        part2, _ =Generate(`[a-Z]{6}`)
        walletID = part1 + part2
      }
      createWalletID(walletID)

      //convert string to date
      if createProfile(request, true, referralID, walletID, referredID) {
        insertEmailMap(request.Mobileno, request.EmailID)
        c.JSON(200, gin.H{
          "status" : "success",
          "first_time_login": false,
          "password_set" : true,
          "profile_set" : true,
          "verified" : true,
        })
      }
    }else if registered && !blocked{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number blocked for unsuccessful verify attempts",
      })
    }else{
      c.JSON(200, gin.H{
        "data":map[string]interface{}{
          "expired": expired,
          "authorized": authorized,
        },
        "status":"failed",
      })
    }
  }
}
