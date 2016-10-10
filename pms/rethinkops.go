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
  createProductTypesTable()
}
func createProductTypesTable() {
  fmt.Println("Creating the productTypes table")
  _, err := r.Branch(
    r.DB("mithun").TableList().Contains("productTypes"),
    nil,
    r.DB("mithun").TableCreate("productTypes"),
  ).Run(session)
  checkErr(err)
}
func insertProductType(option_key int, option_name, option_code string, price int)  bool{
  inserr := r.DB("mithun").Table("productTypes").Insert(productType{
    OptionKey : strconv.Itoa(option_key),
    OptionName: option_name,
    OptionCode: option_code,
    Price: price,
    }).Exec(session)
  checkErr(inserr)
  return true
}
func getProductTypes(choice string)  {

}
