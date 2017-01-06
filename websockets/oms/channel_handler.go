package main

import (
  "github.com/gorilla/websocket"
)

type Topic struct{
  Channel string `json:"topic"`
  Event string `json:"event"`
}
type Admin struct{
  conn *websocket.Conn
  send chan map[string]interface{}
}
type AdminChannels struct{
  admins map[*Admin]bool
  register chan *Admin
  unregister chan *Admin
  broadcast chan map[string]interface{}
  content map[string]interface{}
}

func newAdminChannel() *AdminChannels {
	return &AdminChannels{
    admins:    make(map[*Admin]bool),
		register:   make(chan *Admin),
		unregister: make(chan *Admin),
    broadcast:  make(chan map[string]interface{}),
    content:  map[string]interface{}{},
	}
}

func (h *AdminChannels) run() {
	for {
		select {
  		case c := <-h.register:
  			h.admins[c] = true
  			c.send <- h.content
  			break

  		case c := <-h.unregister:
  			_, ok := h.admins[c]
  			if ok {
  				delete(h.admins, c)
  			}
  			break

  		case m := <-h.broadcast:
  			h.content = m
  			h.broadcastMessage()
  			break
		}
	}
}
func (h *AdminChannels) broadcastMessage() {
	for c := range h.admins {
		select {
    case c.send <- h.content:
  			break
  		default:
  			delete(h.admins, c)
		}
	}
}
