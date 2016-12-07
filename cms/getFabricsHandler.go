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
  DiscRate float64 `json:"disc_rate"`
  Description string `json:"description"`
}
func getFabricsFilterHandler(c *gin.Context){
  var data struct{
    Brand []FilterResult `json:"brand"`
    Category []FilterResult `json:"category"`
    Gender []FilterResult `json:"gender"`
    Quality []FilterResult `json:"quality"`
  }

  data.Brand = getAllFabricsFilter("brand")
  data.Category = getAllFabricsFilter("category")
  data.Quality = getAllFabricsFilter("quality")
  data.Gender = getAllFabricsFilter("gender")

  c.JSON(200, gin.H{
    "status" : "success",
    "data" : data,
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
