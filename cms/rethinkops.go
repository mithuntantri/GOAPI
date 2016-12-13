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
  TotalPrice float64 `gorethink:"total_price"`
  Mobileno string `gorethink:"mobileno"`
  VerifiedUser bool `gorethink:"verified_user"`
  Favorites bool `gorethink:"favorites"`
}
type newCartHash struct{
  Hash string `gorethink:"hash"`
  Mobileno string `gorthink:"mobileno"`
  FabricIds string `gorethink:"fabric_ids"`
  DesignHashs string `gorethink:"design_hashs"`
}
type newBlouseHash struct{
  Hash string `gorethink:"id"`
  Mobileno string `gorethink:"mobileno"`
  CheckedOut bool `gorethink:"checked_out"`
  NeckType string `gorethink:"neck_type"`
  Front string `gorethink:"front"`
  Back string `gorethink:"back"`
  Sleeves string `gorethink:"sleeves"`
  BlouseLength string `gorethink:"blouse_length"`
  Opening string `gorethink:"opening"`
  Cut string `gorethink:"cut"`
  Border string `gorethink:"border"`
  BorderPlacement string `gorethink:"border_placemnt"`
  BorderType string `gorethink:"border_type"`
  Piping string `gorethink:"piping"`
  PipingColor string `gorethink:"piping_color"`
  Dori string `gorethink:"dori"`
  BlousePads string `gorethink:"blouse_pads"`
  DisableList string `gorethink:"disable_list"`
  TotalPrice float64 `gorethink:"total_price"`
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
  createCartHashTable()
  createBlouseHashTable()
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
func createCartHashTable() {
  fmt.Println("Creating the cart HashTable table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("cartHash"),
    nil,
    r.DB("mithun").TableCreate("cartHash"),
  ).Run(session)
  checkErr(err)
}
func createBlouseHashTable(){
  fmt.Println("Creating the blouse HashTable table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("blouseHash"),
    nil,
    r.DB("mithun").TableCreate("blouseHash"),
    ).Run(session)
    checkErr(err)
  }
func checkBlouseHashExists(hash string) bool{
  fmt.Println("Checking if blouse hash already exists")
  result, _ := r.DB("mithun").Table("blouseHash").Get(hash).Run(session)
  if result.IsNil(){
    return false
  }
  return true
}
func checkCartHashExists(hash string) bool{
  fmt.Println("Checking if cart hash already exists")
  result, _ := r.DB("mithun").Table("cartHash").Get(hash).Run(session)
  if result.IsNil(){
    return false
  }
  return true
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
    TotalPrice: 699.00,
    VerifiedUser : VerifiedUser,
    Mobileno : mobileno,
    Favorites : false,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func insertNewBlouseHash(hash, mobileno, neck_type string) bool{
  VerifiedUser := true
  if mobileno == ""{
    VerifiedUser = false
  }
  var disable_list string

  front, list1 := getBlouseOptionKey("Front", neck_type)
  back, list2 := getBlouseOptionKey("Back", neck_type)
  sleeves, list3 := getBlouseOptionKey("Sleeves", "all")
  blouselength, list4 := getBlouseOptionKey("Blouse Length", "all")
  opening, list5 := getBlouseOptionKey("Opening", "all")
  cut, list6 := getBlouseOptionKey("Cut", "all")
  border, list7 := getBlouseOptionKey("Border", "add-on")
  border_placemnt, list8 := getBlouseOptionKey("Border Placement", "add-on")
  border_type, list9 := getBlouseOptionKey("Border Type", "add-on")
  piping, list10 := getBlouseOptionKey("Piping", "add-on")
  piping_color, list11 := getBlouseOptionKey("Piping Color", "add-on")
  dori, list12 := getBlouseOptionKey("Dori", "add-on")
  blouse_pads, list13 := getBlouseOptionKey("Blouse Pads", "add-on")

  disable_list = disable_list + "," + list1 + "," + list2 + "," + list3 + "," + list4 + "," + list5 + "," + list6 + "," + list7 + "," + list8 + "," + list9 + "," + list10 + "," + list11 + "," + list12 + "," + list13
  inserr := r.DB("mithun").Table("blouseHash").Insert(newBlouseHash{
    Hash : hash,
    CheckedOut: false,
    NeckType : neck_type,
    Front : front,
    Back : back,
    Sleeves : sleeves,
    BlouseLength : blouselength,
    Opening :  opening,
    Cut: cut,
    Border : border,
    BorderPlacement: border_placemnt,
    BorderType: border_type,
    Piping  : piping,
    PipingColor : piping_color,
    Dori: dori,
    BlousePads : blouse_pads,
    TotalPrice: 699.00,
    VerifiedUser : VerifiedUser,
    Mobileno : mobileno,
    Favorites : false,
    DisableList : disable_list,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func getNewBlouseHash(hash string) initBlouseData{
  var initdata initBlouseData
  var blousehash newBlouseHash
  curr, _ := r.DB("mithun").Table("blouseHash").Get(hash).Run(session)
  curr.One(&blousehash)
  curr.Close()
  initdata.Hash = hash
  initdata.TotalPrice = blousehash.TotalPrice
  initdata.Favorites = blousehash.Favorites
  initdata.Gender = "F"
  initdata.CheckedOut = blousehash.CheckedOut
  initdata.Front = fetchBlouseOptionsFromKey(hash, blousehash.Front,"Front", blousehash.NeckType)
  initdata.Back = fetchBlouseOptionsFromKey(hash, blousehash.Back, "Back", blousehash.NeckType)
  initdata.Sleeves = fetchBlouseOptionsFromKey(hash, blousehash.Sleeves, "Sleeves", "all")
  initdata.BlouseLength = fetchBlouseOptionsFromKey(hash, blousehash.BlouseLength, "Blouse Length", "all")
  initdata.Opening = fetchBlouseOptionsFromKey(hash, blousehash.Opening, "Opening", "all")
  initdata.Cut = fetchBlouseOptionsFromKey(hash, blousehash.Cut, "Cut", "all")
  initdata.AddOn = make([]AddOns,0)
  var addon AddOns
  addon.Border = fetchBlouseOptionsFromKey(hash, blousehash.Border, "Border", "add-on")
  addon.BorderPlacement = fetchBlouseOptionsFromKey(hash, blousehash.BorderPlacement, "Border Placement", "add-on")
  addon.BorderTypes = fetchBlouseOptionsFromKey(hash, blousehash.BorderType, "Border Type", "add-on")
  addon.Piping = fetchBlouseOptionsFromKey(hash, blousehash.Piping, "Piping", "add-on")
  addon.PipingColor = fetchBlouseOptionsFromKey(hash, blousehash.PipingColor, "Piping Color", "add-on")
  addon.Dori = fetchBlouseOptionsFromKey(hash, blousehash.Dori, "Dori", "add-on")
  addon.BlousePads = fetchBlouseOptionsFromKey(hash, blousehash.BlousePads, "Blouse Pads", "add-on")
  initdata.AddOn = append(initdata.AddOn, addon)
  return initdata
}
func getDisableList(hash string) string{
  var blousehash newBlouseHash
  curr, _ := r.DB("mithun").Table("blouseHash").Get(hash).Run(session)
  curr.One(&blousehash)
  curr.Close()
  return blousehash.DisableList
}
func getDesignPrice(hash string) float64{
  var design newDesignHash
  curr, _ := r.DB("mithun").Table("designHash").Get(hash).Run(session)
  curr.One(&design)
  curr.Close()
  return design.TotalPrice
}
func insertNewCartHash(hash, mobileno string) bool{
  var item []Items
  inserr := r.DB("mithun").Table("cartHash").Insert(cartResponse{
    Hash : hash,
    Mobileno : mobileno,
    Items : item,
    Coupons : "",
    BagTotal: 0,
    CreditsApplied :0,
    EstimatedVAT : 0,
    CouponDiscount : 0,
    DeliveryCharges : 0,
    OrderTotal : 0,
  }).Exec(session)
  checkErr(inserr)
  return true
}
func getCartHash(mobileno string) string{
  var cart_response cartResponse
  curr, _ := r.DB("mithun").Table("cartHash").Filter(map[string]interface{}{
    "mobileno": mobileno,
    }).Run(session)
  curr.One(&cart_response)
  curr.Close()
  return cart_response.Hash
}
func getCartDetails(hash string) cartResponse{
  var cart_response cartResponse
  curr, _ := r.DB("mithun").Table("cartHash").Get(hash).Run(session)
  curr.One(&cart_response)
  curr.Close()
  return cart_response
}
func addItemtoCart(hash string, item Items, price float64) int{
  var count int
  var cart_response cartResponse
  curr, _ := r.DB("mithun").Table("cartHash").Get(hash).Run(session)
  curr.One(&cart_response)
  curr.Close()
  cart_response.Items = append(cart_response.Items, item)
  count = len(cart_response.Items)
  cart_response.BagTotal = cart_response.BagTotal + price
  cart_response.EstimatedVAT = cart_response.BagTotal * 14 / 100
  cart_response.OrderTotal = cart_response.OrderTotal + cart_response.EstimatedVAT + price
  r.DB("mithun").Table("cartHash").Get(hash).Update(cart_response).Exec(session)
  return count
}
func applyCoupontoCart(hash, coupon string, mobile_device bool) bool{
  var cart_response cartResponse
  curr, _ := r.DB("mithun").Table("cartHash").Get(hash).Run(session)
  curr.One(&cart_response)
  curr.Close()
  firstorder := isFirstOrder(cart_response.Mobileno)
  valid, discount := checkCouponValidity(coupon, mobile_device, firstorder, cart_response.BagTotal)
  if valid{
    cart_response.Coupons = coupon
    cart_response.CouponDiscount = discount
    cart_response.EstimatedVAT = (cart_response.BagTotal - cart_response.CouponDiscount - cart_response.CreditsApplied) * 14 / 100
    cart_response.OrderTotal = cart_response.EstimatedVAT + cart_response.BagTotal - cart_response.CouponDiscount - cart_response.CreditsApplied
    r.DB("mithun").Table("cartHash").Get(hash).Update(cart_response).Exec(session)
    return true
  }
  return false
}
func applyCreditstoCart(hash string) bool{
  var cart_response cartResponse
  curr, _ := r.DB("mithun").Table("cartHash").Get(hash).Run(session)
  curr.One(&cart_response)
  curr.Close()
  if(cart_response.Mobileno == ""){
    return false
  }else{
    wallet_id := getwalletIDFromMobile(cart_response.Mobileno)
    _, _, _, credits := getWallet(wallet_id)
    if(credits <= 0){
      return false
    }
    cart_response.CreditsApplied = credits * 30 / 100
    cart_response.EstimatedVAT = (cart_response.BagTotal - cart_response.CouponDiscount - cart_response.CreditsApplied) * 14 / 100
    cart_response.OrderTotal = cart_response.EstimatedVAT + cart_response.BagTotal - cart_response.CouponDiscount - cart_response.CreditsApplied
    r.DB("mithun").Table("cartHash").Get(hash).Update(cart_response).Exec(session)
    useWallet(wallet_id, cart_response.CreditsApplied)
    return true
  }
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
func updateBlouseHashTable(hash , option_name, option_category, option_type string) bool{
  key := getBlouseOptionKeyFromName(option_name, option_category, option_type)
  var blouseHash newBlouseHash
  curr, _ := r.DB("mithun").Table("blouseHash").Get(hash).Run(session)
  curr.One(&blouseHash)
  curr.Close()
  switch option_category{
    case "Front" : blouseHash.Front = key
    case "Back" : blouseHash.Back = key
    case "Sleeves" : blouseHash.Sleeves = key
    case "Blouse Length" : blouseHash.BlouseLength = key
    case "Opening" : blouseHash.Opening = key
    case "Cut" : blouseHash.Cut = key
  }
  r.DB("mithun").Table("blouseHash").Get(hash).Update(blouseHash).Exec(session)
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
    TotalPrice: 700.00,
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
