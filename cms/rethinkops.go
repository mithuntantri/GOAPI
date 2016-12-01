package main

import (
  "fmt"
  "strconv"
  r "gopkg.in/dancannon/gorethink.v2"
)
var (
	session *r.Session
)
type newRegistration struct {
  Id        string  `gorethink:"id,omitempty"`
  Mobileno  string  `gorethink:"mobileno"`
  EmailID   string  `gorethink:"email_id"`
  Password  string  `gorethink:"password"`
  Verified  bool    `gorethink:"verified"`
  Blocked   bool    `gorethink:"is_blocked"`
  ReferredID  string    `gorethink:"referred_id"`
  ClientID string `gorethink:"client_id"`
  FirstName string `gorethink:"firstname"`
  LastName string `gorethink:"lastname"`
  FBID string `gorthink:"fb_id"`
  Gender string `gorethink:"gender"`
}
type loginTokens struct {
  ID string `gorethink:"id, omitempty"`
  ClientID string `gorethink:"client_id"`
  Token string `gorethink:"token"`
  MobileToken string `gorethink:"mobile_token"`
  EmailToken string `gorethink:"email_token"`
}
type emailTokens struct{
  ID string `gorethink:"id, omitempty"`
  EmailID string `gorethink:"email_id"`
  ClientID string `gorethink:"client_id"`
  Token string `gorethink:"token"`
}
type productType struct{
  OptionKey string `gorethink:"id"`
  OptionName string `gorethink:"option_name"`
  OptionCode string `gorethink:"option_code"`
  Price int `gorethink:"price"`
}
type newDesignHash struct{
  Hash string `gorethink:"id"`
  CheckedOut bool `gorethink:"checked_out"`
  Fit string `gorethink:"fit"`
  Sleeve string `gorethink:"sleeve"`
  Collar string `gorethink:"collar"`
  Cuff string `gorethink:"cuff"`
  Placket string `gorethink:"placket"`
  PocketPlacement string `gorethink:"pocket_placement"`
  PocketType string `gorethink:"pocket_type"`
  PocketLid string `gorethink:"pocket_lid"`
  BackDetails string `gorethink:"back_details"`
  BottomCut string `gorethink:"bottom_cut"`
  TotalPrice string `gorethink:"total_price"`
  Mobileno string `gorethink:"mobileno"`
  VerifiedUser bool `gorethink:"verified_user"`
  Favorites bool `gorethink:"favorites"`
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
  createSalttable()
  createDesignHashTable()
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
func createSalttable() {
  fmt.Println("Creating the salt table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("newSalts"),
    nil,
    r.DB("mithun").TableCreate("newSalts"),
  ).Run(session)
  checkErr(err)
}
func createDesignHashTable() {
  fmt.Println("Creating the design HashTable table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("designHash"),
    nil,
    r.DB("mithun").TableCreate("designHash"),
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
func getRegistrationDetails(mobileno string) (string, string, string, string, string, string, string){
  fmt.Println("Fetching Details to permenantly add to DB")
  curr, _ := r.DB("mithun").Table("newRegistrations").Get(mobileno).Run(session)
  var n newRegistration
  curr.One(&n)
  curr.Close()
  return n.EmailID, n.ClientID, n.Password, n.FBID, n.FirstName, n.LastName, n.Gender
}
func getReferredID(mobileno string) string{
  curr, _ := r.DB("mithun").Table("newRegistrations").Get(mobileno).Run(session)
  var n newRegistration
  curr.One(&n)
  curr.Close()
  return n.ReferredID
}
//handles request for new otp
func createRegistration(mobileno, email_id, password, client_id, referred_id, fb_id, firstname, lastname, gender string) bool{
  inserr := r.DB("mithun").Table("newRegistrations").Insert(newRegistration{
    Id: mobileno,
    Mobileno: mobileno,
    Verified: false,
    Blocked : false,
    ReferredID : referred_id,
    ClientID : client_id,
    Password : password,
    EmailID : email_id,
    FBID : fb_id,
    FirstName : firstname,
    LastName : lastname,
    Gender : gender,
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
func checkTokenExists(id, client_id string, mobile_device bool) bool{
  fmt.Println("Checking if user already exists")
  result, _ := r.DB("mithun").Table("loginTokens").Get(id).Run(session)
  if !result.IsNil() {
    var n loginTokens
    result.One(&n)
    result.Close()
    clientid := n.ClientID
    fmt.Println("checking:", clientid, client_id)
    if client_id == clientid {
      if (mobile_device && n.MobileToken != "") || (n.Token != ""){
        fmt.Println("Verified token")
        return true
      }
    }
  }
  return false
}
func checkTokenEntryExists(id, client_id string, mobile_device bool) bool{
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
func checkSaltExists(id string) bool{
  fmt.Println("Checking if salt already exists")
  result, _ := r.DB("mithun").Table("newSalts").Get(id).Run(session)
  if !result.IsNil() {
    return true
  }
  return false
}
func createApplicationSalt(id, salt string) bool{
  fmt.Println("Creating New Salt Token")
  if checkSaltExists(id){
    return false
  }
  inserr := r.DB("mithun").Table("newSalts").Insert(map[string]interface{}{
    "id": id,
    "salt": salt,
    }).Exec(session)
    checkErr(inserr)
  return true
}
func getApplicationSalt(id string)  string{
  fmt.Println("Verifying Token")
  curr, _ := r.DB("mithun").Table("newSalts").Get(id).Run(session)
  var n loginTokens
  curr.One(&n)
  curr.Close()
  return n.Token
}
//handles request for new otp
func createNewToken(id, client_id, token string, mobile_device bool) bool{
  fmt.Println("Creating New Login Token")
  var mobileToken, webToken, emailToken string
  if mobile_device {
    mobileToken = token
    webToken = ""
  }else{
    mobileToken = ""
    webToken = token
  }
  emailToken = ""
  if checkTokenEntryExists(id,client_id, mobile_device){
    updateToken(id, token, mobile_device)
    return true
  }
  inserr := r.DB("mithun").Table("loginTokens").Insert(loginTokens{
    ID: id,
    ClientID: client_id,
    Token : webToken,
    MobileToken :  mobileToken,
    EmailToken : emailToken,
    }).Exec(session)
    checkErr(inserr)
  return true
}
func verifyToken(id, tokenString string, mobile_device bool) bool{
  fmt.Println("Verifying Token")
  curr, _ := r.DB("mithun").Table("loginTokens").Get(id).Run(session)
  var n loginTokens
  curr.One(&n)
  curr.Close()
  var token string
  if mobile_device{
    token = n.MobileToken
  }else{
    token = n.Token
  }
  if token == tokenString{
    return true
  }else{
    return false
  }
}
func updateToken(id, token string, mobile_device bool) {
  fmt.Println("Updating Token")
  if mobile_device{
    r.DB("mithun").Table("loginTokens").Get(id).Update(map[string]interface{}{
      "mobile_token" : token,
    }).Exec(session)
  }else{
    r.DB("mithun").Table("loginTokens").Get(id).Update(map[string]interface{}{
      "token" : token,
    }).Exec(session)
  }
}
func updateEmailToken(id, token string) {
  fmt.Println("Updating Token", id, token)
  r.DB("mithun").Table("loginTokens").Get(id).Update(map[string]interface{}{
    "email_token" : token,
  }).Exec(session)
}
func verifyEmailToken(id, tokenString string) bool{
  fmt.Println("Verifying Email Token")
  curr, _ := r.DB("mithun").Table("loginTokens").Get(id).Run(session)
  var n loginTokens
  curr.One(&n)
  curr.Close()
  var token string
  token = n.EmailToken
  if token == tokenString{
    return true
  }else{
    return false
  }
}
func deleteToken(id string, mobile_device bool)  {
  fmt.Println("Deleting Token")
  if mobile_device{
    r.DB("mithun").Table("loginTokens").Get(id).Update(map[string]interface{}{
      "mobile_token" : "",
    }).Exec(session)
  }else{
    r.DB("mithun").Table("loginTokens").Get(id).Update(map[string]interface{}{
      "token" : "",
    }).Exec(session)
  }
}
func insertNewHash(hash, mobileno string)  bool{
  VerifiedUser := true
  if mobileno == ""{
    VerifiedUser = false
  }
  inserr := r.DB("mithun").Table("designHash").Insert(newDesignHash{
    Hash : hash,
    CheckedOut: false,
    Fit: "101",
    Sleeve: "201",
    Collar: "301",
    Cuff: "401",
    Placket:"501",
    PocketPlacement:"601",
    PocketType:"701",
    PocketLid:"801",
    BackDetails:"901",
    BottomCut: "1001",
    TotalPrice: "700.00",
    VerifiedUser : VerifiedUser,
    Mobileno : mobileno,
    Favorites : false,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func getDesignHash(hash string, choice , options_count int) (bool,bool){
  var key string
  if options_count >= 10{
    var second_part = strconv.Itoa(options_count)
    key = strconv.Itoa(choice) + second_part
  }else{
    var second_part = "0" + strconv.Itoa(options_count)
    key = strconv.Itoa(choice) + second_part
  }
  var designTemp newDesignHash
  curr, _ := r.DB("mithun").Table("designHash").Get(hash).Run(session)
  curr.One(&designTemp)
  curr.Close()
  selected := false
  switch(choice){
        case 1: if designTemp.Fit == key{
                  selected = true
                }
        case 2: if designTemp.Sleeve == key{
                  selected = true
                }
        case 3: if designTemp.Collar == key{
                  selected = true
                }
        case 4: if designTemp.Cuff == key{
                  selected = true
                }
        case 5: if designTemp.Placket == key{
                  selected = true
                }
        case 6: if designTemp.PocketPlacement == key{
                  selected = true
                }
        case 7: if designTemp.PocketType == key{
                  selected = true
                }
        case 8: if designTemp.PocketLid == key{
                  selected = true
                }
        case 9: if designTemp.BackDetails == key{
                  selected = true
                }
        case 10: if designTemp.BottomCut == key{
                  selected = true
                }
  }
  return selected, designTemp.Favorites
}
func checkoutHash(hash string){
  r.DB("mithun").Table("designHash").Get(hash).Update(newDesignHash{
    CheckedOut: true,
  }).Exec(session)
}
func addtoFav(hash string) bool{
  r.DB("mithun").Table("designHash").Get(hash).Update(map[string]interface{}{
    "favorites" : true,
  }).Exec(session)
  return true
}
func removefromFav(hash string) bool{
  r.DB("mithun").Table("designHash").Get(hash).Update(map[string]interface{}{
    "favorites" : false,
  }).Exec(session)
  return true
}
func updateHashTable(hash string, choice, option_key int)  {
  var designTemp = newDesignHash{
    Hash : hash,
    CheckedOut: false,
    Fit: "101",
    Sleeve: "201",
    Collar: "301",
    Cuff: "401",
    Placket:"501",
    PocketPlacement:"601",
    PocketType:"701",
    PocketLid: "801",
    BackDetails:"901",
    BottomCut: "1001",
    TotalPrice: "700.00",
    }
  var key string
  if(option_key >= 10){
    key = strconv.Itoa(choice) + strconv.Itoa(option_key)
  }else{
    key = strconv.Itoa(choice) + "0" + strconv.Itoa(option_key)
  }
  switch choice{
    case 1 : designTemp.Fit = key
    case 2 : designTemp.Sleeve = key
    case 3 : designTemp.Collar = key
    case 4 : designTemp.Cuff = key
    case 5 : designTemp.Placket = key
    case 6 : designTemp.PocketPlacement = key
    case 7 : designTemp.PocketType = key
    case 8 : designTemp.PocketLid = key
    case 9 : designTemp.BackDetails = key
    case 10 : designTemp.BottomCut = key
  }
  fmt.Println(designTemp.Hash, designTemp.Fit)
  r.DB("mithun").Table("designHash").Get(hash).Update(designTemp).Exec(session)
}
