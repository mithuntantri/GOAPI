package main

import (
  "fmt"
  "database/sql"
  "strings"
  _ "github.com/lib/pq"
)

type FilterResult struct{
  Name string `json:"name"`
  Applied bool `json:"applied"`
}
func getAllFabricsFilter(column_name string) []FilterResult{
  Result := make([]FilterResult, 0)
  rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT %s FROM fabrics",QuoteIdentifier(column_name)))
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var name string
    if err := rows.Scan(&name);err != nil{
      checkErr(err)
    }else{
      var result FilterResult
      result.Name = name
      result.Applied = false
      Result  = append(Result, result)
    }
  }
  return Result
}
func getAllFabrics() []Fabrics{
  Result := make([]Fabrics, 0)
  rows, err := db.Query("SELECT fabric_id, gender, brand, category, material, quality, img1, img2, quantity, rate, disc_rate, title, description FROM fabrics")
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var f Fabrics
    if err := rows.Scan(&f.FabricID, &f.Gender, &f.Brand, &f.Category, &f.Material, &f.Quality, &f.Img1, &f.Img2, &f.Quantity, &f.Rate, &f.DiscRate, &f.Title, &f.Description);err != nil{
      checkErr(err)
    }else{
      Result  = append(Result, f)
    }
  }
  return Result
}
func getFabricPrice(fabric_id string) float64 {
  var price float64
  db.QueryRow("SELECT disc_rate from fabrics WHERE fabric_id=$1", fabric_id).Scan(&price)
  return price
}
func getFilteredFabrics(brand, gender, category, quality, material string, apply_brand, apply_gender, apply_category, apply_quality, apply_material bool) []Fabrics{
  Result := make([]Fabrics, 0)
  number := 0
  all_brands := strings.Split(brand, ",")
  all_categories := strings.Split(category, ",")
  all_qualities := strings.Split(quality, ",")
  all_genders := strings.Split(gender, ",")
  all_materials := strings.Split(material, ",")

  var rows *sql.Rows
  var err error

  statement := "SELECT fabric_id, gender, brand, category, material, quality, img1, img2, quantity, rate, disc_rate, title, description FROM fabrics WHERE"
  if apply_brand{
    number++
    for i:=0;i<=len(all_brands)-1;i++{
      if i==0{
        statement = statement + " (brand='"+all_brands[i]+"'"
      }else{
        statement = statement + " OR brand='"+all_brands[i]+"'"
      }
      if i==len(all_brands)-1{
        statement = statement + ")"
      }
    }
  }
  if apply_category{
    if number != 0{
      statement = statement + " AND "
    }
    number++
    for i:=0;i<=len(all_categories)-1;i++{
      if i==0{
        statement = statement + " (category='"+all_categories[i]+"'"
      }else{
        statement = statement + " OR category='"+all_categories[i]+"'"
      }
      if(i==len(all_categories)-1){
        statement = statement + ")"
      }
    }
  }
  if(apply_quality){
    if number != 0{
      statement = statement + " AND "
    }
    number++
    for i:=0;i<=len(all_qualities)-1;i++{
      if i==0{
        statement = statement + " (quality='"+all_qualities[i]+"'"
      }else{
        statement = statement + " OR quality='"+all_qualities[i]+"'"
      }
      if(i==len(all_qualities)-1){
        statement = statement + ")"
      }
    }
  }
  if(apply_gender){
    if number != 0{
      statement = statement + " AND "
    }
    number++
    for i:=0;i<=len(all_genders)-1;i++{
      if i==0{
        statement = statement + " (gender='"+all_genders[i]+"'"
      }else{
        statement = statement + " OR gender='"+all_genders[i]+"'"
      }
      if(i==len(all_genders)-1){
        statement = statement + ")"
      }
    }
  }
  if(apply_material){
    if number != 0{
      statement = statement + " AND "
    }
    number++
    for i:=0;i<=len(all_materials)-1;i++{
      if i==0{
        statement = statement + " (material='"+all_materials[i]+"'"
      }else{
        statement = statement + " OR material='"+all_materials[i]+"'"
      }
      if(i==len(all_materials)-1){
        statement = statement + ")"
      }
    }
  }
  fmt.Println(statement)
  rows, err = db.Query(statement)
  if err != nil{
    checkErr(err)
  }
  for rows.Next(){
    var f Fabrics
    if err := rows.Scan(&f.FabricID,&f.Gender, &f.Brand, &f.Category, &f.Material, &f.Quality, &f.Img1, &f.Img2, &f.Quantity, &f.Rate, &f.DiscRate, &f.Title, &f.Description);err != nil{
      checkErr(err)
    }else{
      Result  = append(Result, f)
    }
  }
  return Result
}
