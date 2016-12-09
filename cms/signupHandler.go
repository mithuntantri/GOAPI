package main

import (
  "github.com/gin-gonic/gin"
)

func signupHandler(c *gin.Context)  {
  var request struct{
    EmailID     string `json:"email_id"`
    Password    string `json:"password"`
    Mobileno    string `json:"mobileno"`
    ClientID    string `json:"client_id"`
    ReferralID  string `json:"referral_id"`
    FirstName   string `json:"firstname"`
    LastName    string `json:"lastname"`
    FBID        string `json:"fb_id"`
    Gender      string `json:"gender"`
  }
  if c.Bind(&request) == nil{
    isNewMobileNo := false
    isNewEmailID := false
    isValidRefCode := false
    response := false

    if isNewMobileNo = checkNewMobileno(request.Mobileno); isNewMobileNo{
      if isNewEmailID = checkNewEmailID(request.EmailID); isNewEmailID{
        // Check if the user already attempted Registration
        registered, blocked, verified := checkRegistrationExists(request.Mobileno)
        if !blocked && !verified{
          //Checking if the referral Id is valid
          if isValidRefCode = checkReferralID(request.ReferralID); isValidRefCode {
            //Update referral_count
            //Call OTP Server
            if response = callnewOTP(request.EmailID, request.Mobileno, "n"); response{
              //Create an Entry in New registrations
              hashedPass := bcryptPassword(request.Password)
              createRegistration(request.Mobileno, request.EmailID, hashedPass, request.ClientID, request.ReferralID, request.FBID, request.FirstName, request.LastName, request.Gender)
            }
            c.JSON(200, gin.H{
              "status" : "success",
              "message": "",
              "data":map[string]interface{}{
                "is_new_mobile": isNewMobileNo,
                "is_new_email": isNewEmailID,
                "is_valid_refcode": isValidRefCode,
                "otp_generated" : response,
              },
            })
          }
        }else if registered && blocked{
          c.JSON(200, gin.H{
            "status" : "failed",
            "message" : "Mobileno blocked because of unsuccessful verify attempts",
            "data":map[string]interface{}{},
          })
        }else if registered && verified{
          c.JSON(200, gin.H{
            "status" : "failed",
            "message" : "Mobile Number already verified",
            "data":map[string]interface{}{},
          })
        }
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Email ID already taken",
        })
      }
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Mobile Number already Registred",
      })
    }
  }
}

// package main
//
// import (
//   "github.com/gin-gonic/gin"
// )
//
// func signupHandler(c *gin.Context)  {
//   var request struct {
//     Mobileno    string `json:"mobileno"`
//     ClientID    string `json:"client_id"`
//     ReferralID  string `json:"referral_id"`
//     EmailID string `json:"email_id"`
//     Password string `json:"password"`
//   }
//   if c.Bind(&request) == nil {
//     isNewUser := false
//     isValidRefCode := false
//     response := false
//     //Checking if the user already registered
//     if isNewUser = checkNewUser(request.Mobileno); isNewUser {
//       // Check if the user already attempted Registration
//       registered, blocked, verified := checkRegistrationExists(request.Mobileno)
//       if !blocked && !verified{
//         //Checking if the referral Id is valid
//         if isValidRefCode = checkReferralID(request.ReferralID); isValidRefCode {
//           //Update referral_count
//           //Call OTP Server
//           if response = callnewOTP(request.Mobileno, "n"); response{
//             //Create an Entry in New registrations
//             createRegistration(request.Mobileno, request.ClientID, request.ReferralID)
//           }
//         }
//         c.JSON(200, gin.H{
//           "status" : "success",
//           "message": "",
//           "data":map[string]interface{}{
//             "is_new_user": isNewUser,
//             "is_valid_refcode": isValidRefCode,
//             "otp_generated" : response,
//           },
//         })
//       }else if registered && blocked{
//         c.JSON(200, gin.H{
//           "status" : "failed",
//           "message" : "Mobileno blocked because of unsuccessful verify attempts",
//           "data":map[string]interface{}{},
//         })
//       }else if registered && verified{
//         c.JSON(200, gin.H{
//           "status" : "failed",
//           "message" : "Mobile Number already verified",
//           "data":map[string]interface{}{},
//         })
//       }
//     }else{
//       c.JSON(200, gin.H{
//         "status" : "success",
//         "message": "",
//         "data":map[string]interface{}{
//           "is_new_user": isNewUser,
//           "is_valid_refcode" : isValidRefCode,
//           "otp_generated" : false,
//         },
//       })
//     }
//   }
// }
