package main

import (
  "github.com/gin-gonic/gin"
)

var fbSecretKey = []byte("4a060ae8248e89896f2999d12a503478")

func main()  {
    connectPSQL()
    connectDB()
    router := gin.Default()
    router.POST("/api/signup", signupHandler)
    router.POST("/api/signup/verify", verifySignupHandler)
    router.POST("/api/signup/resend", resendSignupHandler)
    router.POST("/api/signup/setpasword", setpaswordHandler)
    router.POST("/api/signup/setprofile", setprofileHandler)
    router.POST("/api/login", loginHandler)
    router.POST("/api/login/otp", loginwithotpHanlder)
    router.POST("/api/login/facebook", loginwithfbHandler)
    router.POST("/api/logout", logoutHandler)
    router.Run(":4444")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
