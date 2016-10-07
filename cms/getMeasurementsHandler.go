package main

import (
  "github.com/gin-gonic/gin"
)
type measurements struct {
  MeasurementID string `json:"measurement_id"`
  Mobileno string `json:"mobileno"`
  Units string `json:"units"`
  Neck  string `json:"neck"`
  Chest string `json:"chest"`
  Waist string `json:"waist"`
  Hip string `json:"hip"`
  Length string `json:"length"`
  Shoulder string `json:"shoulder"`
  Sleeve string `json:"sleeve"`
}
func getMeasurementsHandler(c *gin.Context)  {
  var request struct {
    Mobileno string `form:"mobileno" binding:"required"`
    ClientID string `form:"client_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    tokenString := c.Request.Header.Get("X-Authorization-Token")
    device := c.Request.Header.Get("Device")
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
      measurement_id := getMeasurementsID(request.Mobileno)
      measurements := getMeasurements(measurement_id)
      c.JSON(200, gin.H{
        "status" : "success",
        "data" : measurements,
      })
    }
  }
}
