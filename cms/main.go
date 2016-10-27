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

    router.POST("/api/login", loginHandler)
    router.DELETE("/api/logout", logoutHandler)

    router.GET("/api/profile", getProfileHandler)
    router.POST("/api/profile", setprofileHandler)

    router.GET("/api/coupons", getCouponsHandler)

    router.GET("/api/address", getAddressesHandler)
    router.POST("/api/address", createAddressHandler)
    router.PUT("/api/address", updateAddressHandler)
    router.DELETE("/api/address", deleteAddressHandler)

    router.GET("/api/measurements", getMeasurementsHandler)
    router.POST("/api/measurements", createMeasurementsHandler)
    router.PUT("/api/measurements", updateMeasurementsHandler)
    router.DELETE("/api/measurements", deleteMeasurementsHandler)

    router.GET("/api/orders", getOrdersHandler)
    router.POST("/api/orders", createOrdersHandler)

    router.GET("/api/slots", getSlotsHandler)

    router.Run(":3333")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
