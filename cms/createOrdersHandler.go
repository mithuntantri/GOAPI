package main

import (
  "time"
  "strings"
  "github.com/gin-gonic/gin"
)
type order struct{
  OrderID string `json:"order_id"`
  OrderTime int64 `json:"order_time"`
  Mobileno string `json:"mobileno"`
  DesignHash string `json:"design_hash"`
  ProductId string `json:"product_id"`
  MeasurementID string `json:"measurement_id"`
  FabricPickup bool `json:"fabric_pickup"`
  MeasurementPickUp bool `json:"measurement_pickup"`
  AddressID string `json:"address_id"`
  CouponID string `json:"coupon_id"`
  Credits float32 `json:"credits"`
  Discount float32 `json:"discount"`
  FinalPrice float32 `json:"final_price"`
}
func createOrdersHandler(c *gin.Context)  {
  var request order
  if c.Bind(&request) == nil{
    request.OrderTime = time.Now().Unix()
    //generate an order ID
    first := strings.SplitN(request.Mobileno,"", 5)
    part1 := strings.ToUpper(first[0] + first[1] + first[2] + first[3])
    part2, _ := Generate(`[a-Z]{6}`)
    request.OrderID = part1 + part2
    if checkOrderID(request.OrderID) {
      part2, _ = Generate(`[a-Z]{6}`)
      request.OrderID = part1 + part2
    }
    created := createOrder(request)
    if created{
      c.JSON(200, gin.H{
        "status" : "success",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "error",
      })
    }
  }
}
