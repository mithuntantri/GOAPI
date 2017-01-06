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
  Gender    string `json:"gender"`
}
func setprofileHandler(c *gin.Context)  {
  var request profileRequest
  if c.Bind(&request) == nil {
    tokenString := c.Request.Header.Get("X-Authorization-Token")
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    if tokenString == "" {
      c.JSON(401, gin.H{
        "status" : "failed",
        "message" : "Invalid token",
        "data":map[string]interface{}{},
      })
      return
    }
    expired, authorized := authenticateToken(request.Mobileno, request.ClientID, tokenString, mobile_device)
    registered, blocked, verified := checkRegistrationExists(request.Mobileno)
    if !expired && authorized && registered && verified{
      //Get Refferred ID
      referredID := getReferredID(request.Mobileno)

      //Generate a Referral ID
      first := strings.SplitN(request.Mobileno,"", 5)
      part1 := strings.ToUpper(first[0] + first[1] + first[2] + first[3])
      part2, _ := Generate(`[a-Z]{6}`)
      referralID := part1 + part2
      if checkReferralID(referralID) {
        part2, _ =Generate(`[a-Z]{6}`)
        referralID = part1 + part2
      }

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
      createReferralID(referralID, walletID)

      //convert string to date
      if createProfile(request, referralID, walletID, referredID) {
        insertEmailMap(request.Mobileno, request.EmailID)
        updateWallet(walletID, "profile_credits")
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "",
          "data":map[string]interface{}{
            "first_time_login": false,
            "password_set" : true,
            "profile_set" : true,
            "verified" : true,
          },
        })
      }
    }else if registered && blocked{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number blocked for unsuccessful verify attempts",
        "data":map[string]interface{}{},
      })
    }else{
      c.JSON(200, gin.H{
        "status":"failed",
        "message" : "",
        "data":map[string]interface{}{
          "expired": expired,
          "authorized": authorized,
        },
      })
    }
  }
}
