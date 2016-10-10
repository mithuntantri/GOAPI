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
  Price int `json:"price"`
}
type newDesignHash struct{
  Hash string `gorethink:"id"`
  Fit string `gorethink:"fit"`
  Sleeve string `gorethink:"sleeve"`
  Collar string `gorethink:"collar"`
  Cuff string `gorethink:"cuff"`
  Placket string `gorethink:"placket"`
  PocketPlacement string `gorethink:"pocket_placement"`
  PocketType string `gorethink:"pocket_type"`
  PocketLid string `gorethink:"pocket_lid"`
  BackDetails string `gorethink:"back_details"`
  Bottom Cut string `gorethink:"bottom_cut"`
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
  inserr := r.DB("mithun").Table("productTypes").Insert(newDesignHash{
    Hash : hash,
    Fit: "101",
    Sleeve: "201",
    Collar: "301",
    Cuff: "401"
    Placket:"501",
    Pocket Placement:"601",
    Pocket Type:"701",
    Pocket Lid:"No",
    Back Details:"801",
    Bottom Cut: "901",
    }).Exec(session)
  checkErr(inserr)
  return true
}
func updateHashTable(hash, choice, option_key string)  {

}
