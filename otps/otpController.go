package main

import (
  "fmt"
  "strings"
  "encoding/base32"
  "time"
  "github.com/hgfischer/go-otp"
)
func expireOTPchannel(mobileno string) {
    timeChan := time.NewTimer(60 * time.Second).C
    fmt.Println("OTP WILL EXPIRE IN 60sec")
    for {
        select {
        case <- timeChan:
            deleteOTP(mobileno)
            fmt.Println("OTP EXPIRED for",mobileno)
      }
    }
}
func generate_otp(mobileno string) string{
  var (
  	secret   = mobileno
  	isBase32 = false
  	length   = otp.DefaultLength
  	period   = otp.DefaultPeriod
  )
  fmt.Println("Generating OTP using HOTP/TOTP standards")

  key := secret
  if !isBase32 {
    key = base32.StdEncoding.EncodeToString([]byte(secret))
  }

  key = strings.ToUpper(key)

  if !isGoogleAuthenticatorCompatible(key) {
    fmt.Println("WARN: Google Authenticator requires 16 chars base32 secret, without padding")
  }

  fmt.Println("Secret Base32 Encoded Key: ", key)

  totp := &otp.TOTP{
    Secret:         key,
    Length:         uint8(length),
    Period:         uint8(period),
    IsBase32Secret: true,
  }

  otp := totp.Get()
  fmt.Println("TOTP:", otp)
  return otp
}
func isGoogleAuthenticatorCompatible(base32Secret string) bool {
	cleaned := strings.Replace(base32Secret, "=", "", -1)
	cleaned = strings.Replace(cleaned, " ", "", -1)
	return len(cleaned) == 16
}
