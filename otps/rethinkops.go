package main

import (
  "fmt"
  r "gopkg.in/dancannon/gorethink.v2"
)
var (
	session *r.Session
)
type resendUpdate struct{
  ResendAttempts int8 `gorethink:"resend_attempts"`
}
type verifyUpdate struct{
  VerifyAttempts int8 `gorethink:"verify_attempts"`
  Verified bool `gorethink:"verified"`
}
type newOtp struct {
  Id        string  `gorethink:"id,omitempty"`
  Mobileno  string  `gorethink:"mobileno"`
  Otp       string  `gorethink:"otp"`
  Verified  bool    `gorethink:"verified"`
  VerifyAttempts  int8    `gorethink:"verify_attempts"`
  ResendAttempts  int8    `gorethink:"resend_attempts"`
}
func connectDB()  {
  var err error
  session, err = r.Connect(r.ConnectOpts{
    Address:  "localhost:28015",
    Database: "mithun",
    MaxOpen:  40,
  })
  checkErr(err)
  createDB()
}
func createDB()  {
  fmt.Println("Creating the Database")
  _, err := r.Branch(
    r.DBList().Contains("mithun"),
    nil,
    r.DBCreate("mithun"),
  ).Run(session)
  checkErr(err)
  createOTPtable()
}
func createOTPtable() {
  fmt.Println("Creating the newOtp table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("newOtp"),
    nil,
    r.DB("mithun").TableCreate("newOtp"),
  ).Run(session)
  checkErr(err)
}
func checkExists(mobileno string) bool{
  fmt.Println("Checking if user already exists")
  result, _ := r.DB("mithun").Table("newOtp").Get(mobileno).Run(session)
  if result.IsNil() {
    return false
  }
  return true
}
//handles request for new otp
func createOTP(mobileno, otp string) bool{
  inserr := r.DB("mithun").Table("newOtp").Insert(newOtp{
    Id: mobileno,
    Mobileno: mobileno,
    Otp: otp,
    Verified: false,
    VerifyAttempts: 0,
    ResendAttempts: 0,
    }).Exec(session)
  checkErr(inserr)
  return true
}
//handles request for resend otp
func resendOTP(mobileno string) bool{
  fmt.Println("incresing attempts")
  curr, _ := r.DB("mithun").Table("newOtp").Get(mobileno).Run(session)
  var n newOtp
  curr.One(&n)
  curr.Close()
  curr_attempts := n.ResendAttempts
  curr_attempts++
  fmt.Println("Attempts",curr_attempts)
  if curr_attempts <= 3{
    r.DB("mithun").Table("newOtp").Get(mobileno).Update(resendUpdate{
      ResendAttempts: curr_attempts,
      }).Exec(session)
    return false
  }
  return true
}
//handles request for new otp
func verifyOTP(mobileno, otp string) (bool,bool){
  fmt.Println("incresing attempts")
  curr, _ := r.DB("mithun").Table("newOtp").Get(mobileno).Run(session)
  var n newOtp
  curr.One(&n)
  curr.Close()
  curr_attempts := n.VerifyAttempts
  curr_attempts++
  fmt.Println("Attempts",curr_attempts)
  if curr_attempts <= 3{
    db_otp := n.Otp
    if otp == db_otp{
      r.DB("mithun").Table("newOtp").Get(mobileno).Update(verifyUpdate{
        VerifyAttempts: curr_attempts,
        Verified: true,
      }).Exec(session)
      return true, false
    }else{
      r.DB("mithun").Table("newOtp").Get(mobileno).Update(verifyUpdate{
        VerifyAttempts: curr_attempts,
        Verified: false,
      }).Exec(session)
      return false, false
    }
  }
  return false, true
}
func deleteOTP(mobileno string)  {
  fmt.Println("Deleting OTP")
  r.DB("mithun").Table("newOtp").Filter(map[string]interface{}{
    "mobileno": mobileno,
  }).Delete().Exec(session)
}
