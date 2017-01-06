package main

import (
  "github.com/gin-gonic/gin"
)
func updateMeasurementsHandler(c *gin.Context)  {
  var request measurements
  if c.Bind(&request) == nil{
    tokenString := c.Request.Header.Get("X-Authorization-Token")
    device := c.Request.Header.Get("X-Device-Type")
    mobile_device := false
    if device == "mobile"{
      mobile_device = true
    }
    if tokenString == "" {
      c.JSON(401, gin.H{
        "status" : "error",
        "message" : "Request Unauthorized",
      })
      return
    }
    expired, authorized := authenticateToken(request.Mobileno, request.ClientID, tokenString, mobile_device)
    if expired || !authorized{
      c.JSON(401, gin.H{
        "status": "error",
        "message": "Request Unauthorized",
      })
    }else{
        updated := updateMeasurements(request)
        if updated {
          c.JSON(200, gin.H{
            "status": "success",
            "message":"Measurements Updated Successfully",
          })
        }else{
          c.JSON(200, gin.H{
            "status": "success",
            "message":"Failed to Update Measurements",
          })
        }
    }
  }
}
