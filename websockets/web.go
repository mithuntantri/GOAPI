package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)
func main()  {
  go SendName()
  go SendTicker()
  router := gin.Default()
  router.GET("/gowebsocket/websocket", websocketHandler)
  router.Run(":5900")
}
