package main

import (
  "github.com/gin-gonic/gin"
)

func updateProfileHandler(c *gin.Context)  {
  var request profileRequest
  if c.Bind(&request) == nil{
    updateProfile(request)
    c.JSON(200, gin.H{
      "status" : "success",
      "message"  : "Profile Updated Successfully",
    })
  }
}
