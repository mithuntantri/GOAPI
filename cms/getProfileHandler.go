package main

import (
  "github.com/gin-gonic/gin"
)

func getProfileHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `form:"mobileno" binding:"required"`
    ClientID string `form:"client_id" binding:"required"`
  }

  if c.Bind(&request) == nil {
    tokenString := c.Request.Header.Get("X-Authorization-Token")
    device := c.Request.Header.Get("Device")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    if tokenString == "" {
      c.JSON(401, gin.H{
        "status" : "error",
        "message" : "Request Unauthorized",
      })
      return
    }
    expired, authorized := authenticateToken(request.Mobileno, request.ClientID, tokenString, mobile_device)
    if expired || !authorized{
      c.JSON(401, gin.H{
        "status": "error",
        "message": "Request Unauthorized",
      })
    }else{
      response, referral_id, wallet_id := getProfile(request.Mobileno)
      referral_credits, profile_credits, promo_credits := getWallet(wallet_id)
      if response.Mobileno == "" {
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "No Profile Exists",
        })
        return
      }
      c.JSON(200, gin.H{
        "status" : "success",
        "data" : map[string]interface{}{
          "referral_id" : referral_id,
          "personal_info" : map[string]interface{}{
            "mobileno" : response.Mobileno,
            "email_id" : response.EmailID,
            "first_name" : response.FirstName,
            "last_name" : response.LastName,
            "gender" : response.Gender,
          },
          "saved_address" : map[string]interface{}{
            "address" : "",
            "street": "",
            "pin_code" : "",
          },
          "saved_bank": map[string]interface{}{},
          "saved_cards": map[string]interface{}{},
          "subscriptions": map[string]interface{}{},
          "credits": map[string]interface{}{
            "referral_credits" : referral_credits,
            "profile_credits" : profile_credits,
            "promo_credits" : promo_credits,
          },
        },
      })
    }
  }
}
