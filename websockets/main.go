package main

import (
   "net/http"
  //  "io"
   "github.com/gin-gonic/gin"
   "github.com/gorilla/websocket"
  //  "golang.org/x/net/websocket"
        "fmt"
        "time"
)

// setup a channel for delivering stocks
var stockChannel = make(chan int64, 5)
var newChannel = make(chan string, 10)

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type request struct{
  ConnectionType string `json:"conn_type"`
  ConnectionChan string `json:"conn_chan"`
}
func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    for {
        _, msg, _ := conn.ReadMessage()
        fmt.Println("msg", string(msg[:]))
        for{
          if string(msg[:]) == "close" {
            fmt.Println("Closing Connection")
            err := conn.Close()
            if err != nil {
                fmt.Println("Failed to close websocket", err)
                return
            }
          }else if string(msg[:]) == "c:n"{
            data:= <-newChannel
            var response = map[string]interface{}{
              "ref" : "1",
              "data" : data,
            }
            conn.WriteJSON(response)
          }else{
            data:= <- stockChannel
            var response = map[string]interface{}{
              "ref" : "1",
              "data" : data,
            }
            conn.WriteJSON(response)
          }
        }
    }
}

func SendTicker() {
    var tick int64 = 0
    for {
        stockChannel <- tick
        tick += 1
        time.Sleep(3*1e9)
    }
}
func SendName()  {
    var tick int64 = 0
    name := "Mithun"
    for{
      newChannel <- name
      tick += 1
      time.Sleep(3*1e9)
    }
}
func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
        wshandler(c.Writer, c.Request)
    })
  go SendName()
  go SendTicker()
  router.Run(":5900")
  // http.Handle("/", websocket.Handler(StockTickersServer));
  // err := http.ListenAndServe("127.0.0.1:5900", nil);
  // if err != nil {
      // panic(err)
  // }
}
