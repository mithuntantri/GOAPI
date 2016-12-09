package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)
type Items struct{
  ItemHash string `json:"item_hash" gorethink:"item"`
  Quantity float64 `json:"quantity" gorethink:"quantity"`
  Size string `json:"size" gorethink:"size"`
  MeasurementID string `json:"measuremnet_id" gorethink:"measuremnet_id"`
  FabricPickup bool `json:"fabric_pickup" gorethink:"fabric_pickup"`
  MeasurementPickip bool `json:"measurement_pickup" gorethink:"measurement_pickup"`
}
type cartResponse struct{
  Hash string `json:"hash" gorethink:"id"`
  Mobileno string `json:"mobileno" gorethink:"mobileno"`
  Items []Items `json:"items" gorethink:"items"`
  Coupons string `json:"coupons" gorethink:"coupons"`
  BagTotal float64 `json:"bag_total" gorethink:"bag_total"`
  CreditsApplied float64 `json:"credits_applied" gorethink:"credits_applied"`
  EstimatedVAT float64 `json:"estimated_vat" gorethink:"estimated_vat"`
  CouponDiscount float64 `json:"coupon_discount" gorethink:"coupon_discount"`
  DeliveryCharges float64 `json:"delivery_charges" gorethink:"delivery_charges"`
  OrderTotal float64 `json:"order_total" gorethink:"order_total"`
}
func getCartHandler(c *gin.Context)  {
  var request struct{
    Hash string `form:"hash"`
    Mobileno string `form:"mobileno"`
  }
  if c.Bind(&request) == nil{
    if request.Mobileno != "" && request.Hash != ""{
      //get carthash for mobileno
      if checkCartHashExists(request.Hash){
        var cart_response cartResponse
        cart_response = getCartDetails(request.Hash)
        c.JSON(200,gin.H{
          "status" : "success",
          "data" : cart_response,
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Failed to fetch Cart Details",
        })
      }
    }else if request.Mobileno != "" && request.Hash == ""{
      hash := getCartHash(request.Mobileno)
      fmt.Println(hash)
      if hash != ""{
        fmt.Println("Hash already Exists.. Getting details")
        var cart_response cartResponse
        cart_response = getCartDetails(hash)
        c.JSON(200,gin.H{
          "status" : "success",
          "data" : cart_response,
        })
      }else{
        fmt.Println("Hash doesnot exist.. creating details")
        hash,_ := Generate(`[a-Z]{20}`)
        insertNewCartHash(hash, request.Mobileno)
        var cart_response cartResponse
        cart_response = getCartDetails(hash)
        c.JSON(200,gin.H{
          "status" : "success",
          "data" : cart_response,
        })
      }
    }else if(request.Hash != "" && request.Mobileno == ""){
      if checkCartHashExists(request.Hash){
        var cart_response cartResponse
        cart_response = getCartDetails(request.Hash)
        c.JSON(200,gin.H{
          "status" : "success",
          "data" : cart_response,
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Failed to fetch Cart Details",
        })
      }
    }else{
      hash,_ := Generate(`[a-Z]{20}`)
      mobileno := ""
      insertNewCartHash(hash, mobileno)
      var cart_response cartResponse
      cart_response = getCartDetails(hash)
      c.JSON(200,gin.H{
        "status" : "success",
        "data" : cart_response,
      })
    }
  }
}
