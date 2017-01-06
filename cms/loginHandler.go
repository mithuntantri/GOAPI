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
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    var mobileno string
    // var hashedPass string
    var logintoken loginTokens
    if request.Password != ""{
      // hashedPass = getHashedPassword(request.ID, request.Password)
    }
    credentialsExist := false
    ID_is_email := false

    if credentialsExist = checkCredentials(request.ID, request.ClientID, request.Password, true);!credentialsExist{
      credentialsExist = checkCredentials(request.ID, request.ClientID, request.Password, false)
      ID_is_email = credentialsExist
    }
    if !credentialsExist && !request.OtpLogin{
      if !checkIfEmailID(request.ID){
        if !checkIfUsername(request.ID){
          isFblogin := checkIfFBID(request.ID)
          if !isFblogin && !request.FBLogin {
            c.JSON(200, gin.H{
              "status" : "failed",
              "message" : "Credentials does not exist",
              "data":map[string]interface{}{
                "validUser" : false,
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
        }
      }else{
        //Credentials exists with Emailid
        mobileno = getMobileNumberFromEmail(request.ID)
      }
      if checkCredentials(mobileno, request.ClientID, request.Password, true){
        fmt.Println("Genrating login token", mobileno)
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
        if response := callnewOTP("",request.ID, "n"); response{
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
      fmt.Println("New token generating")
      if ID_is_email{
          mobileno = getMobileNumberFromEmail(request.ID)
          logintoken = generateToken(mobileno, request.ClientID, request.Expiry)
      }else{
        logintoken = generateToken(request.ID, request.ClientID, request.Expiry)
      }
    }
    isExists := checkTokenExists(logintoken.ID, logintoken.ClientID, mobile_device)
    if !isExists {
      inserr := createNewToken(logintoken.ID, logintoken.ClientID, logintoken.Token, mobile_device)
      if logintoken.Token == ""{
        c.JSON(200, gin.H{
          "status":"failed",
          "message":"Invalid UserID or Password",
          "data":map[string]interface{}{
              "validUser": true,
              "secret":logintoken.Token,
          },
        })
      }else if inserr {
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
