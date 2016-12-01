package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)
type Fabrics struct{
  FabricID string `json:"fabric_id"`
  Brand string `json:"brand"`
  Category string `json:"category"`
  Quality string `json:"quality"`
  Img string `json:"img"`
  Quantity float64 `json:"quantity"`
  Rate float64 `json:"rate"`
}
func getFabricsFilterHandler(c *gin.Context){
  var brand, category, quality Result
  brand = getAllFabricsFilter("brand")
  category = getAllFabricsFilter("category")
  quality = getAllFabricsFilter("quality")
  c.JSON(200, gin.H{
    "status" : "success",
    "data" : map[string]interface{}{
      "brand" : brand,
      "category" : category,
      "quality" : quality,
    },
  })
}
func getFabricsHandler(c *gin.Context)  {
  var request struct{
    Brand string `form:"brand"`
    Category string `form:"category"`
    Quality string `form:"qualty"`
  }
  if c.Bind(&request) == nil{
    apply_brand := request.Brand != ""
    apply_category := request.Category != ""
    apply_quality := request.Quality != ""
    fmt.Println(apply_brand)
    var data []Fabrics
    if apply_brand || apply_category || apply_quality{
      data = getFilteredFabrics(request.Brand, request.Category, request.Quality, apply_brand, apply_category, apply_quality)
    }else{
      data = getAllFabrics()
    }
    c.JSON(200, gin.H{
      "status" : "success",
      "data" : data,
    })
  }
}
