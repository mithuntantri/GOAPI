package main

import (
  "github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context)  {
  tokenString := c.Request.Header.Get("X-Authorization-Token")
  result := deleteauthToken(tokenString)
  c.JSON(200, gin.H{
    "status" : result,
  })
}
