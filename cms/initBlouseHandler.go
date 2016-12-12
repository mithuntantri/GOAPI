package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)
type blouseOptions struct{
    Name string `json:"name"`
    Category string `json:"category"`
    Type string `json:"type"`
    Img string `json:"img"`
    Enabled bool `json:"enabled"`
    Selected bool `json:"selected"`
    Price string `json:"price"`
}
type initBlouseData struct{
  Hash string `json:"hash"`
  TotalPrice string `json:"total_price"`
  Favorites bool `json:"favorites"`
  Gender string `json:"gender"`
  CheckedOut bool `json:"cheked_out"`
  Front []blouseOptions `json:"front"`
  Back []blouseOptions `json:"back"`
  Sleeves []blouseOptions `json:"sleeves"`
  BlouseLength []blouseOptions `json:"blouse_length"`
  Opening []blouseOptions `json:"opening"`
  Cut []blouseOptions `json:"cut"`
}
func initBlouseHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno"`
    NeckType string `form:"neck_type"`
  }
  if c.Bind(&request) == nil{
    if request.NeckType == ""{
      var response []blouseOptions
      response = fetchBlouseOptions("Neck Type", "all")
      c.JSON(200, gin.H{
        "status" : "success",
        "data" : response,
      })
    }else{
      var initdata initBlouseData
      initdata.Hash, _ = Generate(`[a-Z]{20}`)
      initdata.TotalPrice = "699.00"
      initdata.Favorites = false
      initdata.Gender = "F"
      initdata.CheckedOut = false
      initData.Front = fetchBlouseOptions("Front", request.NeckType)
      initData.Back = fetchBlouseOptions("Back", request.NeckType)
      initData.Sleeves = fetchBlouseOptions("Sleeves", "all")
      initData.BlouseLength = fetchBlouseOptions("Blouse Length", "all")
      initData.Opening = fetchBlouseOptions("Opening", "all")
      initData.Cut = fetchBlouseOptions("Cut", "all")
      insertNewBlouseHash(initData.Hash, request.Mobileno, request.NeckType)
      c.JSON(200, gin.H{
        "status": "success",
        "data": initdata,
      })
    }
  }
}
