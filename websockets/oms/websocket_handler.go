package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
  ReadBufferSize : 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
      return true
  },
}
func websocketHandler(c *gin.Context)  {
  var request struct{
    Id string `form:"id" binding:"required"`
    ClientID string `form:"client_id" binding:"required"`
    Token string `form:"token" binding:"required"`
  }
  if c.Bind(&request) == nil{
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err !=  nil{
      fmt.Println("Failed to set websocket upgrade: %+v",err)
      return
    }
    admin := &Admin{
      conn: conn,
      send: make(chan map[string]interface{}, 256),
    }
    fmt.Println("Websocket Connected Successfully")
    go admin.listenRequests()
    go admin.writeResponse()
  }
}
func (c *Admin) writeResponse(){
	for {
		select {
		  case message, ok := <-c.send:
            if ok{
                c.conn.WriteJSON(message)
            }
    }
  }
}
