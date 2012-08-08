package main

import (
	"github.com/bketelsen/skynet"
	"github.com/bketelsen/skynet/client"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
)

func NewInstanceSocket(ws *websocket.Conn, im *client.InstanceMonitor){
  l := im.Listen(skynet.UUID(), &client.Query{})

  // TODO: make sure this goes out of scope when the user closes the socket or times out (send heartbeat?)
  // Close the websocket, and remove the listener from the InstanceMonitor: l.Close()
  for {
    select {
      case service := <-l.AddChan:
        b, _ := json.Marshal(service)

        ws.Write(b)
      case path := <-l.RemoveChan:
        ws.Write([]byte(path))
    }
  }
}
