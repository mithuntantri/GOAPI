package main

import (
  "github.com/gin-gonic/gin"
)
func addtoFavoritesHandler(c *gin.Context)  {
  var request struct{
    Hash string `json:"hash"`
  }
  if c.Bind(&request) == nil{
    addtoFav(request.Hash)
    c.JSON(200, gin.H{
      "status" : "success",
      "message" : "Added to Favorites Successfully",
    })
  }
}
