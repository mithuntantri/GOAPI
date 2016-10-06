package main

import (
  "github.com/gin-gonic/gin"
)

var mySigningKey = []byte("appsigningsecret")

var fbSecretKey = []byte("4a060ae8248e89896f2999d12a503478")

func main()  {
    connectPSQL()
    connectDB()
    router := gin.Default()
    router.POST("/api/signup", signupHandler)
    router.POST("/api/signup/verify", verifySignupHandler)
    router.POST("/api/signup/resend", resendSignupHandler)
    router.POST("/api/setpassword", setpaswordHandler)
    router.POST("/api/setprofile", setprofileHandler)
    router.POST("/api/login", loginHandler)
    router.DELETE("/api/logout", logoutHandler)
    router.POST("/api/admin/add", adminAddHandler)
    router.POST("/api/measurements", updateMeasurementsHandler)
    router.GET("/api/measurements", getMeasurementsHandler)
    router.GET("/api/profile", getProfileHandler)
    router.Run(":3333")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
