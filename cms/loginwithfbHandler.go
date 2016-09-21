package main

import (
  "github.com/gin-gonic/gin"
)

func loginwithfbHandler(c *gin.Context)  {
  var request struct{
    ID string `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
  }
  if c.Bind(&request) == nil {
    exists := checkIfFBID(request.ID)
    c.JSON(200, gin.H{
      "status" : "success",
      "verfied" : exists,
    })
  }
}
