package main

import (
  "github.com/gin-gonic/gin"
)

func checkCouponHandler(c *gin.Context)  {
  var request struct{
    Mobileno string `form:"mobileno"`
    CouponCode string `form:"coupon_code" binding:"required"`
    Amount float64 `form:"amount" binding:"required"`
  }
  if c.Bind(&request) == nil{
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    if request.Mobileno == ""{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Failed to apply Coupon. Please Login/Signup",
      })
    }else{
      //Check if the order is the first order
      firstorder := isFirstOrder(request.Mobileno)
      valid, discount := checkCouponValidity(request.CouponCode, mobile_device, firstorder, request.Amount)
      if valid{
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Valid Coupon Code",
          "data" : map[string]interface{}{
            "discount" : discount,
          },
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "failed",
          "message" : "Invalid Coupon Code",
        })
      }
    }
  }
}
