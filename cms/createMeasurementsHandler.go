package main

import (
  "strings"
  "github.com/gin-gonic/gin"
)
func createMeasurementsHandler(c *gin.Context)  {
    var request measurements
    if c.Bind(&request) == nil{
      //Generate a Measurement ID
      first := strings.SplitN(request.Mobileno,"", 5)
      part1 := strings.ToUpper(first[0] + first[1] + first[2] + first[3])
      part2, _ := Generate(`[a-Z]{6}`)
      measurementID := part1 + part2
      if checkMeasurementID(measurementID) {
        part2, _ =Generate(`[a-Z]{6}`)
        measurementID = part1 + part2
      }
      created := createMeasurementsID(measurementID, request.Mobileno)
      if created{
        c.JSON(200, gin.H{
          "status" : "success",
          "message" : "Measurements Added Successfully",
        })
      }else{
        c.JSON(200, gin.H{
          "status" : "error",
          "messaged": "Failed to add measurements",
        })
      }
    }
}
