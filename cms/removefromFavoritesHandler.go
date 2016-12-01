package main

import (
  "github.com/gin-gonic/gin"
)
func removefromFavoritesHandler(c *gin.Context)  {
  var request struct{
    Hash string `form:"hash" binding:"required"`
  }
  if c.Bind(&request) == nil{
    removefromFav(request.Hash)
    c.JSON(200, gin.H{
      "status" : "success",
      "message" : "Removed from Favorites Successfully",
    })
  }
}
