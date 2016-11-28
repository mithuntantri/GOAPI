package main

import (
  "github.com/gin-gonic/gin"
)
type Items struct{
  ItemHash string `json:"item_hash"`
}
var cartResponse struct{
  Items []Items `json:"items"`
  Coupons string `json:"coupons"`
  BagTotal string `json:"bag_total"`
  CreditsApplied string `json:"credits_applied"`
  EstimatedVAT string `json:"estimated_vat"`
  CouponDiscount string `json:"coupon_discount"`
  DeliveryCharges string `json:"delivery_charges"`
  OrderTotal string `json:"order_total"`
}
func getCartHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno" binding:"required"`
  }
  if c.Bind(&request) == nil{

  }
}
