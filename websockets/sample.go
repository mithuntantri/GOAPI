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
  broadcast chan string
  register chan *client
  unregister chan *client
  content string
}
var chan1 = allChannel{
  broadcast:   make(chan string),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients: 	 make(map[*client]bool),
	content:  	 "Joined Chan1",
}
var chan2 = allChannel{
  broadcast:   make(chan string),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients: 	 make(map[*client]bool),
	content:  	 "Joined Chan2",
}
func (h *allChannel) run() {
	for {
		select {
  		case c := <-h.register:
  			h.clients[c] = true
  			c.send <- []byte(h.content)
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
        fmt.Println("Broadcasting: ",h.content)
  			h.broadcastMessage()
  			break
		}
	}
}
func (h *allChannel) broadcastMessage() {
	for c := range h.clients {
		select {
    case c.send <- []byte(h.content):
      fmt.Println("BroadcastingMessage: ",h.content)
  			break
  		default:
  			close(c.send)
  			delete(h.clients, c)
		}
	}
}
type client struct{
  conn *websocket.Conn
  send chan []byte
  channel chan []byte
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
    send: make(chan []byte, 1024),
    channel: make(chan []byte, 1024),
  }
	go c.writePump()
  go c.requestListener()
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
// setup a channel for delivering stocks
var stockChannel = make(chan string, 5)
var newChannel = make(chan string, 10)
func SendTicker() {
  var tick int64 = 0
  name := "c:d"
  for{
    newChannel <- name
    tick += 1
    time.Sleep(3*1e9)
  }
}
func SendName()  {
    var tick int64 = 0
    name := "c:n"
    for{
      stockChannel <- name
      tick += 1
      time.Sleep(3*1e9)
    }
}
func (c *client) requestListener(){
  var topic topic
  c.conn.ReadJSON(&topic)
  c.channel <- []byte(topic.Channel)
  for {
    select {
    case newRequest := <- c.channel:
      fmt.Println("New Request Received", newRequest)
    }
  }
}
func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
      fmt.Println("Sending message", message)
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *client) write(mt int, message []byte) error {
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, message)
}
func (c *client) readPump()  {
  defer func() {
    chan1.unregister <- c
		chan2.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait));
		return nil
	})

	for {
    var topic topic
		c.conn.ReadJSON(&topic)
    fmt.Println("New connection requets")
    if(topic.Channel == "c:n"){
      fmt.Println("Registering for chan 1")
      chan1.register <- c
      for{
        data := <- stockChannel
        chan1.broadcast <- data
      }
    }else if(topic.Channel == "c:d"){
      fmt.Println("Registering for chan 2")
      chan2.register <- c
      for{
        data := <- newChannel
        chan2.broadcast <- data
      }
    }else{
      // fmt.Println("Waiting for connection..")
      // for{
      //   data := <- stockChannel
      //   chan1.broadcast <- data
      // }
    }
	}
}
func main()  {
  go chan1.run()
  go chan2.run()
  go SendName()
  go SendTicker()
  router := gin.Default()
  router.GET("/gowebsocket/websocket", websocketHandler)
  router.Run(":5900")
}
