package main

import (
  "fmt"
  "crypto/tls"
  "gopkg.in/gomail.v2"
)
func sendOtpThroughMail(email_id, otp string)  {
  fmt.Println("Sending OTP Email")
  mail := gomail.NewMessage()
  mail.SetHeader("From", "noreply@zigfo.com")
  mail.SetHeader("To", email_id)
  mail.SetHeader("Subject", "Verify Mobileno - Access your Zigfo Account Securely")
  mail.SetBody("text/html", `
    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
    <html xmlns="http://www.w3.org/1999/xhtml">

    <head>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title></title>
    </head>

    <body style="font-family:-apple-system, '.SFNSText-Regular', 'Helvetica Neue', Roboto, 'Segoe UI', sans-serif; color: #666666; background:white; text-decoration: none;">
      <table width="100%" cellpadding="0" cellspacing="0" border="0" summary="">
        <tr align="center">
          <td style="height: 30px; width: 100%;">&nbsp;</td>
        </tr>
        <tr align="center">
          <td valign="top" style="width: 100%;">
            <table cellspacing="0" cellpadding="0" border="0" summary="">
              <tr align="center">
                <td valign="middle" style="width: 100%;">
                  <a href="https://www.zigfo.com">
                    <img width="55" height="71" style="border:0;width:55px;height:71px;" src="http://www.zigfo.com/img/mail-logo.png" alt="zigfo">
                  </a>
                </td>
              </tr>
            </table>
          </td>
        </tr>
        <tr align="center">
          <td style="height: 50px; width: 100%;">&nbsp;</td>
        </tr>
        <tr align="center">
          <td valign="top" style="width: 100%;">
            <table style="padding: 0px; border: 0; max-width: 520px; text-align: center;" width="100%" cellpadding="0" cellspacing="0" border="0" summary="">
              <tr align="center">
                <td style="width: 100%; margin: 0px 10px; line-height: 24px; font-size: 14pt; font-weight: bold; color: #333333;">
                  <p style="margin: 0; padding: 0;">Verify your Mobile Number</p>
                </td>
              </tr>
              <tr align="center" style="margin: 0px 10px;">
                <td style="width: 100%; line-height: 24px; font-size: 11pt;">
                  <p>Why? Just one more step to access your Zigfo account. Your One Time Password(OTP) is also sent to your registered Mobile Number.</p>
                </td>
              </tr>
              <tr align="center">
                <td style="height: 30px; width: 100%;">&nbsp;</td>
              </tr>
              <tr align="center">
                <td style="width: 100%; margin: 0px 10px; line-height: 24px; font-size: 11pt;">
                  <a style="padding: 10px 20px; border: 1px solid #1492ef; -webkit-border-radius: 999em; -moz-border-radius: 999em; border-radius: 999em; line-height: 24px; font-size: 11pt; background-color: #1492ef; color: white; text-decoration: none;"
                  href="#">OTP : `+ otp +`</a>
                </td>
              </tr>
              <tr>
                <td style="height: 55px; width: 100%;">&nbsp;</td>
              </tr>
            </table>
            <table style="border-collapse:collapse; max-width: 520px; text-align: center;" cellpadding="0" cellspacing="0" border="0" summary="">
              <tr>
                <td style="height: 50px; width: 100%;">&nbsp;</td>
              </tr>
              <tr align="center">
                <td style="width: 100%;">
                  <p style="line-height: 20px; font-size: 10pt; color: #b3b3b3;">Get the Zigfo app:</p>
                </td>
              </tr>
              <tr align="center">
                <td style="width: 100%;">
                  <a href="https://itunes.apple.com/app/zigfo-social-bookmarks/id1056141950" style="text-decoration:none;">
                    <img width="135" height="40" style="border: 0; width: 135px; height: 40px; margin-left: 0px; margin-right: 3px;" src="http://www.zigfo.com/img/apple-app-store.png" alt="iOS app"
                      title="iOS app">
                  </a>
                  <a href="https://play.google.com/store/apps/details?id=com.zigfo.android" style="text-decoration:none;">
                    <img width="135" height="40" style="border: 0; width: 135px; height: 40px; margin-left: 3px; margin-right: 0px;" src="http://www.zigfo.com/img/google-play-store.png" alt="Android app"
                      title="Android app">
                  </a>
                </td>
              </tr>
              <tr>
                <td style="height: 55px; width: 100%;">&nbsp;</td>
              </tr>
              <tr align="center">
                <td style="width: 100%;">
                  <p style="line-height: 20px; font-size: 9pt; color: #b3b3b3;">Sent by
                    <a href="https://www.zigfo.com" style="color: #1492ef; text-decoration: none;">zigfo</a>
                  </p>
                  <p style="font-size: 9pt; color: #b3b3b3;">
                    <a href="https://itunes.apple.com/app/zigfo" style="text-decoration:none;">
                      <img width="20" height="20" style="border: 0; width: 18px; height: 18px; margin-right:6px;" src="http://www.zigfo.com/img/apple.png" alt="iOS app" title="iOS app">
                    </a>
                    <a href="https://twitter.com/zigfocom" style="text-decoration:none;">
                      <img width="20" height="20" style="border: 0; width: 18px; height: 18px; margin-right:6px;" src="http://www.zigfo.com/img/twitter.png" alt="Twitter" title="Twitter">
                    </a>
                    <a href="https://www.facebook.com/zigfocom" style="text-decoration:none;">
                      <img width="18" height="18" style="border: 0; width: 18px; height: 18px;" src="/http://www.zigfo.com/img/facebook.png" alt="Facebook" title="Facebook">
                    </a>
                  </p>
                </td>
              </tr>
            </table>
          </td>
        </tr>
        <tr align="center">
          <td style="height: 55px; width: 100%;">&nbsp;</td>
        </tr>
      </table>
    </body>

    </html>
  `)
  d := gomail.NewPlainDialer("smtp.zoho.com", 587, "noreply@zigfo.com", "password123")
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
    if err := d.DialAndSend(mail); err != nil {
        panic(err)
    }
}
