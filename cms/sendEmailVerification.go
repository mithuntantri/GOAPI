package main

import (
  "time"
  "fmt"
  "log"
  "crypto/tls"
  "gopkg.in/gomail.v2"
  "github.com/dgrijalva/jwt-go"
)
func generateEmailToken(mobileno, client_id, email_id string, expiry bool)  emailTokens{
  token := jwt.New(jwt.SigningMethodHS256)
  if expiry {
    token.Claims =  jwt.MapClaims{
      "id"  : mobileno,
      "email_id" : email_id,
      "client_id" : client_id,
      "exp"   : time.Now().Add(time.Minute * 10).Unix(),
    }
  }else{
    token.Claims =  jwt.MapClaims{
      "id"  : mobileno,
      "email_id" : email_id,
      "client_id" : client_id,
      "exp"   : time.Now().Add(time.Minute * 100).Unix(),
    }
  }
  tokenString, err := token.SignedString(mySigningKey)
  if(err != nil){
    log.Fatal(err)
  }
  //Strore login tokens for all users
  emailToken := emailTokens{
    ID : mobileno,
    EmailID : email_id,
    ClientID: client_id,
    Token: tokenString,
  }
  return emailToken
}
func sendEmailVerification(mobileno, client_id, email_id string)  {
  var emailToken emailTokens
  emailToken = generateEmailToken(mobileno, client_id, email_id, true)
  fmt.Println("Sending Email", mobileno, email_id, emailToken.Token)
  updateEmailToken(mobileno, emailToken.Token)
  url := `www.zigfo.com/confirm-email?token=` + emailToken.Token
  mail := gomail.NewMessage()
  mail.SetHeader("From", "noreply@zigfo.com")
  mail.SetHeader("To", email_id)
  mail.SetHeader("Subject", "Verify Email - Access your Zigfo Account Securely")
  mail.SetBody("text/html", `
    <table style="max-width: 630px;border-left: solid 1px #e6e6e6;border-right: solid 1px #e6e6e6;width:100%;background-color:#00436d;">
      <td style="width:100%;height:60px;"></td>
    </table>
    <div style="max-width:600px;width:100%;background-color:#fafafa;padding:15px;">
      <div style="line-height:30px;">Hi,</div>
      <div style="line-height:30px;">Greetings!</div>
      <div style="line-height:30px;">Just one more step to access your Zigfo account</div>
      <div style="line-height:30px;">In this email you are provided a Verification Link to access your account.</div>
      <div style="padding-top:10px;padding-bottom:10px;">This is only a temporary link. You'll be prompted to redirected to our website and this is to ensure that only you have access to your account.</div>
      <div style="padding-top:10px;">Your Verification Link: <b><a href="` + url + `">Click Here!</a></b></div>
      <div style="padding-bottom:30px;">Expires in: <b>10mins only</b></div>
      <div>Best Regards, </div>
    <div><b>Team Zigfo</b></div></div>
  `)
  d := gomail.NewPlainDialer("smtp.zoho.com", 587, "noreply@zigfo.com", "password123")
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(mail); err != nil {
        panic(err)
    }
}
