package main

import (
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)
const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type allChannel struct{
  clients map[*client]bool
  broadcast chan map[string]interface{}
  register chan *client
  unregister chan *client
  content map[string]interface{}
}
var chan1 = allChannel{
  broadcast:   make(chan map[string]interface{}),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients: 	 make(map[*client]bool),
	content:  	 map[string]interface{}{},
}
var chan2 = allChannel{
  broadcast:   make(chan map[string]interface{}),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients: 	 make(map[*client]bool),
	content:  	 map[string]interface{}{},
}
func (h *allChannel) run() {
	for {
		select {
  		case c := <-h.register:
  			h.clients[c] = true
  			c.send <- (h.content)
  			break

  		case c := <-h.unregister:
  			_, ok := h.clients[c]
  			if ok {
  				delete(h.clients, c)
  				close(c.send)
  			}
  			break

  		case m := <-h.broadcast:
  			h.content = m
  			h.broadcastMessage()
  			break
		}
	}
}
func (h *allChannel) broadcastMessage() {
	for c := range h.clients {
		select {
    case c.send <- h.content:
  			break
  		default:
  			close(c.send)
  			delete(h.clients, c)
		}
	}
}
type client struct{
  conn *websocket.Conn
  send chan map[string]interface{}
}
var upgrader = websocket.Upgrader{
  ReadBufferSize : 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
      return true
  },
}
type topic struct{
  Channel string `json:"channel"`
  Event string `json:"event"`
}
func wsHandler(w http.ResponseWriter, r *http.Request)  {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil{
    fmt.Println("Failed to set websocket upgrade: %+v", err)
    return
  }
  c := &client{
    conn: conn,
    send: make(chan map[string]interface{}, 1024),
  }

	go c.writePump()
	c.readPump()
}
func websocketHandler(c *gin.Context){
  var request struct{
    Token string `form:"token" binding:"required"`
    LoginID string `form:"login_id" binding:"required"`
  }
  if c.Bind(&request) == nil{
    wsHandler(c.Writer, c.Request)
  }
}
func (c *client) writePump()  {
  fmt.Println("Starting Listen Channel for Client")
  chan1.register <- c
  fmt.Println(chan1.register)
}
func (c *client) readPump()  {

}
func main()  {
  go chan1.run()
  go chan2.run()
  router := gin.Default()
  router.GET("/gowebsocket/websocket", websocketHandler)
  router.Run(":5900")
}
