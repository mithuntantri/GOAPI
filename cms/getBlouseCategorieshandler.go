package main

import (
  "github.com/gin-gonic/gin"
)

func getBlouseCategoriesHandler(c *gin.Context)  {
  var response []blouseOptions
  response = fetchBlouseOptions("", "Neck Type", "all")
  c.JSON(200, gin.H{
    "status" : "success",
    "data" : response,
  })
}
