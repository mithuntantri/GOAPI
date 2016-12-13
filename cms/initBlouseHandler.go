package main

import (
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
    DisableList string `json:"disable_list"`
}
type AddOns struct{
  Border []blouseOptions `json:"border"`
  BorderPlacement []blouseOptions `json:"border_placement"`
  BorderTypes []blouseOptions `json:"border_types"`
  Piping []blouseOptions `json:"piping"`
  PipingColor []blouseOptions `json:"piping_color"`
  Dori []blouseOptions `json:"dori"`
  BlousePads []blouseOptions `json:"blouse_pads"`
}
type initBlouseData struct{
  Hash string `json:"hash"`
  TotalPrice float64 `json:"total_price"`
  Favorites bool `json:"favorites"`
  Gender string `json:"gender"`
  CheckedOut bool `json:"cheked_out"`
  Front []blouseOptions `json:"front"`
  Back []blouseOptions `json:"back"`
  Sleeves []blouseOptions `json:"sleeves"`
  BlouseLength []blouseOptions `json:"blouse_length"`
  Opening []blouseOptions `json:"opening"`
  Cut []blouseOptions `json:"cut"`
  AddOn []AddOns `json:"add-ons"`
}
func initBlouseHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno"`
    NeckType string `form:"neck_type" binding:"required"`
  }
  if c.Bind(&request) == nil{
    var initdata initBlouseData
    initdata.Hash, _ = Generate(`[a-Z]{20}`)
    insertNewBlouseHash(initdata.Hash, request.Mobileno, request.NeckType)
    initdata.TotalPrice = 699.00
    initdata.Favorites = false
    initdata.Gender = "F"
    initdata.CheckedOut = false
    initdata.Front = fetchBlouseOptions(initdata.Hash, "Front", request.NeckType)
    initdata.Back = fetchBlouseOptions(initdata.Hash, "Back", request.NeckType)
    initdata.Sleeves = fetchBlouseOptions(initdata.Hash, "Sleeves", "all")
    initdata.BlouseLength = fetchBlouseOptions(initdata.Hash, "Blouse Length", "all")
    initdata.Opening = fetchBlouseOptions(initdata.Hash, "Opening", "all")
    initdata.Cut = fetchBlouseOptions(initdata.Hash, "Cut", "all")
    initdata.AddOn = make([]AddOns,0)

    var addon AddOns
    addon.Border = fetchBlouseOptions(initdata.Hash, "Border", "add-on")
    addon.BorderPlacement = fetchBlouseOptions(initdata.Hash, "Border Placement", "add-on")
    addon.BorderTypes = fetchBlouseOptions(initdata.Hash, "Border Types", "add-on")
    addon.Piping = fetchBlouseOptions(initdata.Hash, "Piping", "add-on")
    addon.PipingColor = fetchBlouseOptions(initdata.Hash, "Piping Color", "add-on")
    addon.Dori = fetchBlouseOptions(initdata.Hash, "Dori", "add-on")
    addon.BlousePads = fetchBlouseOptions(initdata.Hash, "Blouse Pads", "add-on")
    initdata.AddOn = append(initdata.AddOn, addon)

    c.JSON(200, gin.H{
      "status": "success",
      "data": initdata,
    })
  }
}
