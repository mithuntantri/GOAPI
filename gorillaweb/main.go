package main

import (
  "fmt"
  "flag"
  "github.com/gin-gonic/gin"
  "github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:6789", "http service address")
