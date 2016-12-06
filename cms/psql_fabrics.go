package main

import (
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

type Result struct{
  Name []string `json:"name"`
}
func getAllFabricsFilter(column_name string) Result{
  var Result Result
  Result.Name = make([]string, 0)
  rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT %s FROM fabrics",QuoteIdentifier(column_name)))
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var name string
    if err := rows.Scan(&name);err != nil{
      checkErr(err)
    }else{
      Result.Name  = append(Result.Name, name)
    }
  }
  return Result
}
func getAllFabrics() []Fabrics{
  Result := make([]Fabrics, 0)
  rows, err := db.Query("SELECT fabric_id, gender, brand, category, quality, img, quantity, rate, disc_rate, description FROM fabrics")
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var f Fabrics
    if err := rows.Scan(&f.FabricID, &f.Gender, &f.Brand, &f.Category, &f.Quality, &f.Img, &f.Quantity, &f.Rate, &f.DiscRate, &f.Description);err != nil{
      checkErr(err)
    }else{
      Result  = append(Result, f)
    }
  }
  return Result
}
func getFilteredFabrics(brand, gender, category, quality string, apply_brand, apply_gender, apply_category, apply_quality bool) []Fabrics{
  Result := make([]Fabrics, 0)
  number := 0
  fmt.Println(brand, gender, category, quality, apply_brand, apply_gender, apply_category, apply_quality)
  var column1, column2, column3, column4, column1_value, column2_value, column3_value, column4_value string
  if apply_brand{
    number++
    column1 = "brand"
    column1_value = brand
  }
  if apply_category{
    number++
    if column1 != ""{
      column2 = "category"
      column2_value = category
    }else{
      column1 = "category"
      column1_value = category
    }
  }
  if apply_quality{
    number++
    if column1 == ""{
      column1 = "quality"
      column1_value = quality
    }else if column2 == ""{
      column2 = "quality"
      column2_value = quality
    }else{
      column3 = "quality"
      column3_value = quality
    }
  }
  if apply_gender{
    number++
    if column1 == ""{
      column1 ="gender"
      column1_value = gender
    }else if column2 == ""{
      column2 = "gender"
      column2_value = gender
    }else if column3 == ""{
      column3 = "gender"
      column3_value = gender
    }else{
      column4 = "gender"
      column4_value = gender
    }
  }
  fmt.Println(column1, column2, column3, column4, column1_value, column2_value, column3_value, column4_value)
  var rows *sql.Rows
  var err error
  if number == 1{
    rows, err = db.Query(fmt.Sprintf("SELECT fabric_id, gender, brand, category, quality, img, quantity, rate, disc_rate, description FROM fabrics WHERE %s=$1",QuoteIdentifier(column1)),column1_value)
  }else if number == 2{
    rows, err = db.Query(fmt.Sprintf("SELECT fabric_id, gender, brand, category, quality, img, quantity, rate, disc_rate, description FROM fabrics WHERE %s=$1 AND %s=$2",QuoteIdentifier(column1),QuoteIdentifier(column2)),column1_value,column2_value)
  }else if number == 3{
    rows, err = db.Query(fmt.Sprintf("SELECT fabric_id, gender, brand, category, quality, img, quantity, rate, disc_rate, description FROM fabrics WHERE %s=$1 AND %s=$2 AND %s=$3",QuoteIdentifier(column1),QuoteIdentifier(column2),QuoteIdentifier(column3)),column1_value,column2_value,column3_value)
  }else{
    rows, err = db.Query(fmt.Sprintf("SELECT fabric_id, gender, brand, category, quality, img, quantity, rate, disc_rate, description FROM fabrics WHERE %s=$1 AND %s=$2 AND %s=$3 AND %s=$4",QuoteIdentifier(column1),QuoteIdentifier(column2),QuoteIdentifier(column3), QuoteIdentifier(column4)),column1_value,column2_value,column3_value,column4_value)
  }
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var f Fabrics
    if err := rows.Scan(&f.FabricID,&f.Gender, &f.Brand, &f.Category, &f.Quality, &f.Img, &f.Quantity, &f.Rate, &f.DiscRate, &f.Description);err != nil{
      checkErr(err)
    }else{
      Result  = append(Result, f)
    }
  }
  return Result
}
