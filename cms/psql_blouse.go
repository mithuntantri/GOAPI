package main

import (
  "strings"
  _ "github.com/lib/pq"
)
func fetchBlouseOptionsFromKey(hash, option_key, option_category, option_type string) []blouseOptions{
  var blouse []blouseOptions
  rows,err := db.Query("SELECT option_key, option_name, option_category, option_type, img, price, disable_list from blouse WHERE option_category=$1 AND option_type=$2", option_category, option_type)
  if err != nil{
    checkErr(err)
  }
  var key string
  for rows.Next(){
    var b blouseOptions
    if err := rows.Scan(&key, &b.Name, &b.Category, &b.Type, &b.Img, &b.Price, &b.DisableList); err != nil{
      checkErr(err)
    }else{
      if key == option_key{
        b.Selected = true
      }
      b.Enabled = true
      list := getDisableList(hash)
      disable_list := strings.Split(list, ",")
      for i:=0;i< len(disable_list);i++{
        if disable_list[i] == key{
          b.Enabled = false
        }
      }
      blouse = append(blouse, b)
    }
  }
  return blouse
}
func fetchBlouseOptions(hash, option_category, option_type string) []blouseOptions {
  var blouse []blouseOptions
  rows,err := db.Query("SELECT option_key, option_name, option_category, option_type, img, price, disable_list from blouse WHERE option_category=$1 AND option_type=$2", option_category, option_type)
  if err != nil{
    checkErr(err)
  }
  var key1 string
  key2,_ := getBlouseOptionKey(option_category, option_type)

  for rows.Next(){
    var b blouseOptions
    if err := rows.Scan(&key1, &b.Name, &b.Category, &b.Type, &b.Img, &b.Price, &b.DisableList); err != nil{
      checkErr(err)
    }else{
      if key1 == key2{
        b.Selected = true
      }
      b.Enabled = true
      list := getDisableList(hash)
      disable_list := strings.Split(list, ",")
      for i:=0;i< len(disable_list);i++{
        if disable_list[i] == key1{
          b.Enabled = false
        }
      }
      blouse = append(blouse, b)
    }
  }
  return blouse
}
func getBlouseOptionKey(option_category, option_type string) (string, string){
  var key, disable_list string
  rows, err := db.Query("SELECT option_key, disable_list FROM blouse WHERE option_category=$1 AND option_type=$2", option_category, option_type)
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    if err := rows.Scan(&key, &disable_list); err != nil{
      checkErr(err)
    }else{
      return key, disable_list
    }
  }
  return key, disable_list
}
func getBlouseOptionKeyFromName(option_name, option_category, option_type string) string{
  var key string
  rows, err := db.Query("SELECT option_key FROM blouse WHERE option_name=$1 AND option_category=$2 AND option_type=$3",option_name, option_category, option_type)
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    if err := rows.Scan(&key); err != nil{
      checkErr(err)
    }else{
      return key
    }
  }
  return key
}
