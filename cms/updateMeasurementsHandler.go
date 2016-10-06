package main

import (
  "github.com/gin-gonic/gin"
)
func updateMeasurementsHandler(c *gin.Context)  {
  // var request struct{
  //   Units string `json:"units"`
  //   Neck int  `json:"neck"`
  //   Chest int `json:"chest"`
  //   Waist int `json:"waist"`
  //   Hip int `json:"hip"`
  //   Length  int `json:"length"`
  //   Shoulder  int `json:"shoulder"`
  //   Sleeve  int `json:"sleeve"`
  // }
  // if c.Bind(&request) == nil{
  //   tokenString := c.Request.Header.Get("X-Authorization-Token")
  //   device := c.Request.Header.Get("Device")
  //   mobile_device := false
  //   if device == "mobile"{
  //     mobile_device = true
  //   }
  //   fmt.Println(tokenString)
  //   if tokenString == "" {
  //     c.JSON(401, gin.H{
  //       "status" : "failed",
  //       "message" : "Invalid token",
  //       "data":map[string]interface{}{},
  //     })
  //     return
  //   }
  //   expired, authorized := authenticateToken(request.Mobileno, request.ClientID, tokenString, mobile_device)
  //   if expired || !authorized{
  //     c.JSON(401, gin.H{
  //       "status": "failed",
  //       "message": "Invalid Token",
  //     })
  //   }else{
  //
  //   }
  // }
}
