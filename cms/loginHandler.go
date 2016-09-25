package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func loginHandler(c *gin.Context)  {
  var request struct {
    ID string `json:"id"`
    ClientID string `json:"client_id"`
    Expiry bool `json:"set_expiry"`
    Password string `json:"password"`
    OtpLogin  bool  `json:"otp_login"`
    ResendOtp bool `json:"resend_otp"`
    Otp string `json:"otp"`
    FBLogin bool  `json:"fb_login"`
  }
  if c.Bind(&request) == nil{
    device := c.Request.Header.Get("Device")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    var mobileno string
    var hashedPass string
    var logintoken loginTokens
    if request.Password != ""{
      hashedPass = getHashedPassword(request.Password)
    }
    credentialsExist := checkCredentials(request.ID, request.ClientID, hashedPass)
    if !credentialsExist && !request.OtpLogin{
      fmt.Println("User exists")
      if !checkIfEmailID(request.ID){
        fmt.Println("User with email id exists")
        if !checkIfUsername(request.ID){
          fmt.Println("User with username id exists")
          isFblogin := !checkIfFBID(request.ID)
          if !isFblogin && request.FBLogin {
            c.JSON(200, gin.H{
              "status" : "success",
              "message" : "",
              "data":map[string]interface{}{
                "valid" : false,
              },
            })
            return
          }else if request.FBLogin && isFblogin{
            //Credentials exists with FB ID
            mobileno = getMobileNumber(request.ID, "fbid_map", "fb_id")
          }
        }else{
          //Credentials Exists with Username
          mobileno = getMobileNumber(request.ID, "username_map", "username")
          fmt.Println(mobileno)
        }
      }else{
        //Credentials exists with Emailid
        mobileno = getMobileNumber(request.ID, "emailid_map", "email_id")
        fmt.Println(mobileno)
      }
      if checkCredentials(mobileno, request.ClientID, hashedPass){
        logintoken = generateToken(mobileno, request.ClientID, request.Expiry)
      }
    }else if credentialsExist && request.OtpLogin && request.Otp == ""{
      //Generate Otp for mobileno or resend otp
        if request.ResendOtp{
          blocked := callresOTP(request.ID, "r")
          c.JSON(200, gin.H{
            "status" : "success",
            "message" : "",
            "data":map[string]interface{}{
              "blocked" : blocked,
            },
          })
          return
        }
        if response := callnewOTP(request.ID, "n"); response{
          c.JSON(200, gin.H{
            "status" : "success",
            "message" : "",
            "data":map[string]interface{}{
              "validUser" : true,
              "otp_generated" : true,
            },
          })
          return
        }
    }else if credentialsExist && request.OtpLogin && request.Otp != ""{
      ///Verify Mobil Number with otp
      blocked, verified := callverOTP(request.ID, request.Otp, "v")
      if !blocked && verified{
        logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "",
          "data":map[string]interface{}{
            "verified" : verified,
            "blocked" : blocked,
          },
        })
        return
      }
    }else if credentialsExist && !request.OtpLogin && !request.FBLogin {
      logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
    }
    isExists := checkTokenExists(logintoken.ID, logintoken.ClientID, mobile_device)
    if !isExists {
      inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token, mobile_device)
      if inserr {
        c.JSON(200, gin.H{
          "status":"success",
          "message":"",
          "data":map[string]interface{}{
              "validUser": true,
              "secret":logintoken.Token,
          },
        })
      }
    }else{
      updateToken(logintoken.ID, logintoken.Token, mobile_device)
      c.JSON(200, gin.H{
        "status":"success",
        "message":"",
        "data":map[string]interface{}{
            "validUser": true,
            "secret": logintoken.Token,
          },
        })
      }
  }
}
