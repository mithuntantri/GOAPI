package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strings"
)

func verifySignupHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `json:"mobileno"`
    Otp string `json:"otp"`
  }
  registered, blocked, verified := false, false, false
  if c.Bind(&request) == nil {
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    //Check for the new Registrations
    registered, blocked, verified = checkRegistrationExists(request.Mobileno)
    fmt.Println("Checking for new Registration", registered, blocked, verified)
    if registered && !blocked && !verified {
        //Verify OTP
        blocked, verified = callverOTP(request.Mobileno, request.Otp, "v")
        //Update Registrations
        updateRegistrations(request.Mobileno, blocked, verified)
        fmt.Println("Adding to DB", verified)
        if verified{
          ReferredID := getReferredID(request.Mobileno)
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

          walletRefID := updateReferralTable(ReferredID)
          updateWallet(walletRefID, "referral_credits")

          EmailID, ClientID, hashedPass, FBID, FirstName, LastName, Gender := getRegistrationDetails(request.Mobileno)
          var profileRequest = profileRequest{
            Mobileno : request.Mobileno,
            EmailID : EmailID,
            ClientID : ClientID,
            FirstName : FirstName,
            LastName : LastName,
            Gender : Gender,
          }
          if addtoCredentials(request.Mobileno, EmailID, true, false, ClientID, hashedPass){
            logintoken := generateToken(request.Mobileno, ClientID, true)
            inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token, mobile_device)
            if FBID != ""{
              insertFbIDMap(request.Mobileno, FBID)
            }
            createProfile(profileRequest, referralID, walletID, ReferredID)
            go sendEmailVerification(request.Mobileno, ClientID, EmailID)
            c.JSON(200, gin.H{
              "status":"success",
              "message":"",
              "data":map[string]interface{}{
                "validUser": true,
                "password_set" :  inserr,
                "profile_set" : false,
                "first_time_login" :  true,
                "secret":logintoken.Token,
              },
            })
          }
        }else{
          c.JSON(200, gin.H{
            "status":"failed",
            "message":"",
            "data":map[string]interface{}{
              "validUser": true,
              "password_set" :  true,
              "profile_set" : false,
              "first_time_login" :  true,
              "secret":"",
            },
          })
        }
    }else if !registered {
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number not requested for OTP",
        "data":map[string]interface{}{},
      })
    }else if registered && blocked {
      c.JSON(200, gin.H{
        "status": "failed",
        "message": "Mobile Number blocked beacuse of unsuccessful verify attempts",
        "data":map[string]interface{}{},
      })
    }else if registered && verified{
      c.JSON(200, gin.H{
        "status": "failed",
        "message": "Mobile Number already registered & verified",
        "data":map[string]interface{}{},
      })
    }
  }
}
