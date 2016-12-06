package main

import (
  "github.com/gin-gonic/gin"
)

func addtoCartHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `json:"mobileno"`
    FabricID string `json:"fabric_id"`
    Quantity float64 `json:"quantity"`
    CouponCode string `json:"coupon_code"`
  }
  if c.Bind(&request) == nil{

  }
}
