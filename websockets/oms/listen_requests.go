package main

import (
  "fmt"
  "github.com/gorilla/websocket"
  "log"
)

func (c *Admin) listenRequests() {
  var topic Topic
  for {
      err := c.conn.ReadJSON(&topic)
      if err != nil {
  			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
  				log.Printf("error: %v", err)
  			}
  			break
  		}
      if(topic.Channel == "c:n"){
        switch (topic.Event){
          case "join" : fmt.Println("Registering for chan 1")
                        orderChannel.register <- c
                        break
          case "close" : fmt.Println("Unregistering for chan1")
                        orderChannel.unregister <- c
                        break
        }
      }else if(topic.Channel == "c:d"){
        switch (topic.Event){
          case "join" : fmt.Println("Registering for chan 2")
                        statsChannel.register <- c
                        break
          case "close" : fmt.Println("Unregistering for chan2")
                        statsChannel.unregister <- c
                        break
        }
      }else{
        fmt.Println("Heartbeat")
      }
  }
}
func appendorderData(){
  for{
   data := <- newChannel
   orderChannel.broadcast <- map[string]interface{}{
     "topic" : "c:n",
     "event" : data,
     "payload": map[string]interface{}{
      "status": "ok",
      "response" : map[string]interface{}{
        "data" : "neworder",
      },
     },
   }
 }
}
func appendstatsData(){
  for{
   data := <- newChannel
   statsChannel.broadcast <- map[string]interface{}{
     "topic" : "c:d",
     "event" : data,
     "payload": map[string]interface{}{
      "status": "ok",
      "response" : map[string]interface{}{
        "data" : "newstats",
      },
   },
   }
 }
}
