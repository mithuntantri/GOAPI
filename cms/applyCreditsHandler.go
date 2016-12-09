package main

import (
  "github.com/gin-gonic/gin"
)

func applyCreditsHandler(c *gin.Context)  {
  var request struct{
    Hash string `json:"hash"`
  }
  if c.Bind(&request) == nil{
    if applyCreditstoCart(request.Hash){
      c.JSON(200, gin.H{
        "status" : "success",
        "message" : "Credits applied Successfully",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Failed to apply credits",
      })
    }
  }
}
