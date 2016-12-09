package main

import (
  "github.com/gin-gonic/gin"
)

func addtoCartHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `json:"mobileno"`
    Hash string `json:"hash"`
    FabricID string `json:"fabric_id"`
    DesignHash string `json:"designhash"`
    Quantity float64 `json:"quantity"`
    Size string `json:"size"`
    MeasurementID string `json:"measuremnet_id"`
    FabricPickup bool `json:"fabric_pickup"`
    MeasurementPickip bool `json:"measurement_pickup"`
  }
  if c.Bind(&request) == nil{
    var item Items
    var price float64
    if !checkCartHashExists(request.Hash){
      c.JSON(200, gin.H{
        "status" : "error",
        "message" : "Invalid Request",
      })
      return
    }
    if(request.FabricID != ""){
      item.ItemHash = request.FabricID
      price = getFabricPrice(request.FabricID)
    }else if(request.DesignHash != ""){
      item.ItemHash = request.DesignHash
      price = getDesignPrice(request.DesignHash)
    }
    item.FabricPickup = request.FabricPickup
    item.MeasurementPickip = request.MeasurementPickip
    if (!request.FabricPickup){
      item.Size = request.Size
    }
    if(!request.MeasurementPickip){
      item.MeasurementID = request.MeasurementID
    }
    item.Quantity = request.Quantity
    if request.Hash != ""{
      count := addItemtoCart(request.Hash, item, price)
      c.JSON(200, gin.H{
        "status" : "success",
        "item_count" : count,
        "message" : "Item added to Cart Successfully",
      })
    }else if (request.Mobileno != ""){
      hash:= getCartHash(request.Mobileno)
      count := addItemtoCart(hash, item, price)
      c.JSON(200, gin.H{
        "status" : "success",
        "item_count" : count,
        "message" : "Item added to Cart Successfully",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "error",
        "message" : "Failed to add item to Cart",
      })
    }
  }
}
