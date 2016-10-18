package main

import (
  "fmt"
  "strconv"
  r "gopkg.in/dancannon/gorethink.v2"
)
var (
	session *r.Session
)
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
  createDesignHashTable()
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
func insertNewHash(hash string)  bool{
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
    }).Exec(session)
  checkErr(inserr)
  return true
}
func getDesignHash(hash string, choice , options_count int) bool{
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
  return selected
}
func checkoutHash(hash string){
  r.DB("mithun").Table("designHash").Get(hash).Update(newDesignHash{
    CheckedOut: true,
  }).Exec(session)
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
