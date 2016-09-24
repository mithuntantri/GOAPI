package main

import (
  "fmt"
  r "gopkg.in/dancannon/gorethink.v2"
)
var (
	session *r.Session
)
type newRegistration struct {
  Id        string  `gorethink:"id,omitempty"`
  Mobileno  string  `gorethink:"mobileno"`
  Verified  bool    `gorethink:"verified"`
  Blocked   bool    `gorethink:"is_blocked"`
  ReferredID  string    `gorethink:"referred_id"`
  ClientID string `gorethink:"client_id"`
}
type loginTokens struct {
  ID string `gorethink:"id, omitempty"`
  ClientID string `gorethink:"client_id"`
  Token string `gorethink:"token"`
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
  createRegistrationstable()
  createTokenstable()
}
func createRegistrationstable() {
  fmt.Println("Creating the newRegistrations table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("newRegistrations"),
    nil,
    r.DB("mithun").TableCreate("newRegistrations"),
  ).Run(session)
  checkErr(err)
}
func checkRegistrationExists(mobileno string) (bool, bool, bool){
  fmt.Println("Checking if user already exists")
  result, _ := r.DB("mithun").Table("newRegistrations").Get(mobileno).Run(session)
  if result.IsNil() {
    return false, false, false
  }else{
    var n newRegistration
    result.One(&n)
    result.Close()
    return true, n.Blocked, n.Verified
  }
}
func getReferredID(mobileno string) string{
  curr, _ := r.DB("mithun").Table("newRegistrations").Get(mobileno).Run(session)
  var n newRegistration
  curr.One(&n)
  curr.Close()
  return n.ReferredID
}
//handles request for new otp
func createRegistration(mobileno, client_id, referred_id string) bool{
  inserr := r.DB("mithun").Table("newRegistrations").Insert(newRegistration{
    Id: mobileno,
    Mobileno: mobileno,
    Verified: false,
    Blocked : false,
    ReferredID : referred_id,
    ClientID : client_id,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func updateRegistrations(mobileno string, blocked bool, verified bool) bool{
  if verified{
    fmt.Println("increasing referred count")
    curr, _ := r.DB("mithun").Table("newRegistrations").Get(mobileno).Run(session)
    var n newRegistration
    curr.One(&n)
    curr.Close()
    if referred_id := n.ReferredID; referred_id != "null" {
      updateReferralTable(referred_id)
    }
  }
  r.DB("mithun").Table("newRegistrations").Get(mobileno).Update(map[string]interface{}{
    "is_blocked": blocked,
    "verified": verified,
    }).Exec(session)
  return true
}
//Delete once the number is Verified
func deleteRegistration(mobileno string)  {
  fmt.Println("Deleting Entry")
  r.DB("mithun").Table("newRegistrations").Filter(map[string]interface{}{
    "mobileno": mobileno,
  }).Delete().Exec(session)
}
func createTokenstable() {
  fmt.Println("Creating the loginTokens table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("loginTokens"),
    nil,
    r.DB("mithun").TableCreate("loginTokens"),
  ).Run(session)
  checkErr(err)
}
func checkTokenExists(id, client_id string) bool{
  fmt.Println("Checking if user already exists")
  result, _ := r.DB("mithun").Table("loginTokens").Get(id).Run(session)
  if !result.IsNil() {
    var n loginTokens
    result.One(&n)
    result.Close()
    clientid := n.ClientID
    fmt.Println("checking:", clientid, client_id)
    if client_id == clientid {
      fmt.Println("Verified token")
      return true
    }
  }
  return false
}
//handles request for new otp
func createNewToken(id, client_id, token string) bool{
  fmt.Println("Creating New Login Token")
  inserr := r.DB("mithun").Table("loginTokens").Insert(loginTokens{
    ID: id,
    ClientID: client_id,
    Token : token,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func verifyToken(id, tokenString string) bool{
  fmt.Println("Verifying Token")
  curr, _ := r.DB("mithun").Table("loginTokens").Get(id).Run(session)
  var n loginTokens
  curr.One(&n)
  curr.Close()
  token := n.Token
  if token == tokenString{
    return true
  }else{
    return false
  }
}
func deleteToken(id string)  {
  fmt.Println("Deleting Token")
  r.DB("mithun").Table("loginTokens").Filter(map[string]interface{}{
    "id": id,
  }).Delete().Exec(session)
}
