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

    router.GET("/api/verifyemail", verifyEmailIDHandler)
    router.GET("/api/resendemail", resendEmailVerificationHandler)

    router.POST("/api/login", loginHandler)
    router.DELETE("/api/logout", logoutHandler)

    router.GET("/api/facebook", facebookHandler)
    router.GET("/api/google", googleHandler)

    router.GET("/api/profile", getProfileHandler)
    router.POST("/api/profile", setprofileHandler)
    router.PUT("/api/profile", updateProfileHandler)

    router.GET("/api/coupons", getCouponsHandler)
    router.POST("/api/coupons", applyCouponsHandler)
    router.POST("/api/credits", applyCreditsHandler)
    router.GET("/api/coupons/validity", checkCouponHandler)

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

    router.GET("/api/product/categories", getCategoriesHandler)
    router.GET("/api/product/init", initProductHandler)
    router.GET("/api/product/options", getOptionsHandler)
    router.POST("/api/product/options", setValueHandler)

    router.GET("/api/blouse/categories", getBlouseCategoriesHandler)
    router.GET("/api/blouse/init", initBlouseHandler)
    router.GET("/api/blouse/options", getBlouseOptionsHandler)
    router.POST("/api/blouse/options", setBlouseValueHandler)

    router.POST("/api/cart", addtoCartHandler)
    router.GET("/api/cart", getCartHandler)

    router.GET("/api/fabrics/filter", getFabricsFilterHandler)
    router.GET("/api/fabrics", getFabricsHandler)

    router.PUT("/api/favorites", addtoFavoritesHandler)
    router.DELETE("/api/favorites", removefromFavoritesHandler)

    router.Run(":3333")
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
