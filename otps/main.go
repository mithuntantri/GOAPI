package main

import (
  "github.com/gin-gonic/gin"
)
func main()  {
  connectDB()
  router := gin.Default()
  router.POST("/otp/create", createotpHandler)
  router.POST("/otp/verify", verifyotpHandler)
  router.POST("/otp/resend", resendotpHandler)
  router.Run(":2000")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
