package main

import (
  "github.com/gin-gonic/gin"
)
type coupons struct{
  CouponID string `json:"coupon_id"`
  Description string `json:"description"`
  Expiry int64 `json:"expiry"`
  OnlyNew bool `json:"only_new"`
  OnlyFirst bool `json:"only_first"`
  OnlyApp bool `json:"only_app"`
}
type Coupon struct{
  AllCoupons []coupons `json:"all_coupons"`
}
func getCouponsHandler(c *gin.Context)  {
  var all_coupons Coupon
  all_coupons = fetchCoupons()
  c.JSON(200,gin.H{
    "status" : "success",
    "data" : all_coupons,
  })
}
