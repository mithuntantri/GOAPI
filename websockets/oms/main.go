package main

import (
  "time"
  "github.com/gin-gonic/gin"
)

var orderChannel *AdminChannels
var statsChannel *AdminChannels

func main()  {
  orderChannel = newAdminChannel()
  statsChannel = newAdminChannel()

  go orderChannel.run()
  go statsChannel.run()

  go appendorderData()
  go appendstatsData()

  go SendTicker()

  router := gin.Default()
  router.GET("/gowebsocket/websocket", websocketHandler)
  router.Run(":5900")
}

var newChannel = make(chan string, 10)
func SendTicker() {
  var tick int64 = 0
  name := "c:d"
  for{
    newChannel <- name
    tick += 1
    time.Sleep(5*1e9)
  }
}
