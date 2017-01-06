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
type client struct{
  conn *websocket.Conn
  send chan map[string]interface{}
  channel chan []byte
}
type allChannel struct{
  clients     map[*client]bool
  broadcast   chan map[string]interface{}
  register    chan *client
  unregister  chan *client
  content     map[string]interface{}
}
var chan1 = allChannel{
  clients: 	   make(map[*client]bool),
  broadcast:   make(chan map[string]interface{}),
	register:    make(chan *client),
	unregister:  make(chan *client),
	content:  	 map[string]interface{}{
                  "status":"joined",
                  "channel":"c:n",
               },
}
var chan2 = allChannel{
  clients: 	   make(map[*client]bool),
  broadcast:   make(chan map[string]interface{}),
	register:    make(chan *client),
	unregister:  make(chan *client),
	content:  	 map[string]interface{}{
                  "status":"joined",
                  "channel":"c:d",
               },
}
func (h *allChannel) run() {
	for {
		select {
  		case c := <-h.register:
  			h.clients[c] = true
        fmt.Println("Registering Client")
        c.send <- h.content
  			break

  		case c := <-h.unregister:
  			_, ok := h.clients[c]
  			if ok {
          fmt.Println("UnRegistering Client")
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
    case c.send <- (h.content):
      fmt.Println("BroadcastingMessage: ",h.content)
  			break
  		default:
  			close(c.send)
  			delete(h.clients, c)
		}
	}
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
    channel: make(chan []byte, 1024),
  }
  go c.writeMsg()
  c.readMsg()
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
func (c *client) joinChannel(){
  var topic topic
  c.conn.ReadJSON(&c)
  fmt.Println("Request for new Join", topic.Channel)
}
func (c *client) writeMsg(){
  ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <- c.send:
      fmt.Println("Sending message", message, ok)
			// if !ok {
				c.conn.WriteJSON(message)
				return
			// }
			// if err := c.write(websocket.TextMessage, []byte("message")); err != nil {
				// return
			// }
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
func (c *client) readMsg(){
  for {
    fmt.Println("Came here")
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
func main()  {
  go chan1.run()
  go chan2.run()
  go SendName()
  go SendTicker()
  router := gin.Default()
  router.GET("/gowebsocket/websocket", websocketHandler)
  router.Run(":5900")
}
