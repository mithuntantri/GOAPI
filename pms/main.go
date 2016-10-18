package main

import (
  "github.com/gin-gonic/gin"
)
func main()  {
  connectPSQL()
  connectDB()
  router := gin.Default()
  router.GET("/product/categories", getCategoriesHandler)
  router.GET("/product/init", initProductHandler)
  router.GET("/product/options", getOptionsHandler)
  router.POST("/product/options", setValueHandler)
  router.Run(":2222")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
