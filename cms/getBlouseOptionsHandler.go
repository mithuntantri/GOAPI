package main

import (
  "github.com/gin-gonic/gin"
)

func getBlouseOptionsHandler(c *gin.Context)  {
  var request struct{
    Hash string `form:"hash" binding:"required"`
  }
  if c.Bind(&request) == nil{
    if checkBlouseHashExists(request.Hash){
      var initdata initBlouseData
      initdata = getNewBlouseHash(request.Hash)
      c.JSON(200, gin.H{
        "status": "success",
        "data": initdata,
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Invalid Blouse Hash",
      })
    }
  }
}
