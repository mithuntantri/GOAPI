package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func signupHandler(c *gin.Context)  {
  var request struct {
    Mobileno    string `json:"mobileno"`
    ClientID    string `json:"client_id"`
    ReferralID  string `json:"referral_id"`
  }
  if c.Bind(&request) == nil {
    isNewUser := false
    isValidRefCode := false
    response := false
    //Checking if the user already registered
    if isNewUser = checkNewUser(request.Mobileno); isNewUser {
      // Check if the user already attempted Registration
      registered, blocked, verified := checkRegistrationExists(request.Mobileno)
      fmt.Println("new",registered, blocked, verified)
      if !blocked && !verified{
        fmt.Println("new",registered, blocked, verified)
        //Checking if the referral Id is valid
        if isValidRefCode = checkReferralID(request.ReferralID); isValidRefCode {
          //Update referral_count
          updateReferralTable(request.ReferralID)
          //Call OTP Server
          fmt.Println("Calling OTP Server")
          if response = callnewOTP(request.Mobileno, "n"); response{
            //Create an Entry in New registrations
            createRegistration(request.Mobileno, request.ClientID, request.ReferralID)
          }
        }
        c.JSON(200, gin.H{
          "status" : "success",
          "is_new_user": isNewUser,
          "is_valid_refcode": isValidRefCode,
          "otp_generated" : response,
        })
      }else if registered && blocked{
        fmt.Println("blocked",registered, blocked, verified)
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Mobileno blocked because of unsuccessful verify attempts",
        })
      }else if registered && verified{
        fmt.Println("verified",registered, blocked, verified)
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Mobile Number already verified",
        })
      }
    }else{
      c.JSON(200, gin.H{
        "status" : "success",
        "is_new_user": isNewUser,
        "is_valid_refcode" : isValidRefCode,
        "otp_generated" : false,
      })
    }
  }
}
