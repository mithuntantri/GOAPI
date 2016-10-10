package main

import (
  "github.com/gin-gonic/gin"
)
func main()  {
  // createProductTypes()
  router := gin.Default()
  router.GET("/product/categories", getCategoriesHandler)
  router.GET("/product/init", initProductHandler)
  router.POST("/product/getOptions", getOptionsHandler)
  router.POST("/product/setValue", setValueHandler)
  router.Run(":2222")
}
func checkErr(err error)  {
  panic(err)
}
