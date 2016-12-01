package main

import (
  "github.com/gin-gonic/gin"
)
type Fabrics struct{
  FabricID string `json:"fabric_id"`
  Brand string `json:"brand"`
  Gender string `json:"gender"`
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
    Gender string `form:"gender"`
    Category string `form:"category"`
    Quality string `form:"quality"`
  }
  if c.Bind(&request) == nil{
    apply_brand := request.Brand != ""
    apply_gender := request.Gender != ""
    apply_category := request.Category != ""
    apply_quality := request.Quality != ""
    var data []Fabrics
    if apply_brand || apply_category || apply_quality || apply_gender{
      data = getFilteredFabrics(request.Brand, request.Gender, request.Category, request.Quality, apply_brand, apply_gender, apply_category, apply_quality)
    }else{
      data = getAllFabrics()
    }
    c.JSON(200, gin.H{
      "status" : "success",
      "data" : data,
    })
  }
}
