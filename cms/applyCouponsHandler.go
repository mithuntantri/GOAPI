package main

import (
  "github.com/gin-gonic/gin"
)
func applyCouponsHandler(c *gin.Context)  {
  var request struct{
    Hash string `json:"hash"`
    Coupon string `json:"coupon"`
  }
  if c.Bind(&request) == nil{
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    if request.Hash == ""{
      c.JSON(200, gin.H{
        "status" : "error",
        "message" : "Failed to apply Coupon. Please Login/Signup",
      })
    }else{
      if applyCoupontoCart(request.Hash, request.Coupon, mobile_device){
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Coupon applied successfully",
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Failed to apply Coupon",
        })
      }
    }
  }
}
