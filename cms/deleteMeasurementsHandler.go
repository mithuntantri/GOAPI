package main

import (
  "github.com/gin-gonic/gin"
)

func deleteMeasurementsHandler(c *gin.Context)  {
  var request struct{
    MeasurementID string `form:"measurement_id" binding:"required"`
    ClientID string `form:"client_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    if checkMeasurementID(request.MeasurementID) {
      deleted := deleteMeasurements(request.MeasurementID)
      if deleted{
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Measurement deleted Successfully",
        })
        }else{
          c.JSON(200, gin.H{
            "status" : "error",
            "message": "Failed to delete Measurement",
          })
        }
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Invalid Measurement ID",
      })
    }
  }
}
