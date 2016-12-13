package main

import (
  "github.com/gin-gonic/gin"
)

func setBlouseValueHandler(c *gin.Context)  {
  var request struct{
    Hash string `json:"hash"`
    Name string `json:"name"`
    Category string `json:"category"`
    Type string `json:"type"`
  }
  if c.Bind(&request) == nil{
    if checkBlouseHashExists(request.Hash){
      updateBlouseHashTable(request.Hash, request.Name, request.Category, request.Type)
      c.JSON(200, gin.H{
        "status" : "success",
        "message" : "Option Changed Successfully",
      })
    }else{
      c.JSON(200, gin.H{
        "status" : "failed",
        "message" : "Invalid Blouse Hash",
      })
    }
  }
}
