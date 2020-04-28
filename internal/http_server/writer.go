package httpserver

import (
	"fmt"
	"io"
	"log"

	"github.com/evakom/sniffer/internal/models"
	"github.com/gorilla/websocket"
)

type writer struct {
	io.Writer
	conn []*websocket.Conn
}

func (w *writer) Write(p []byte) (int, error) {
	m := models.Message{
		Type: models.MTMessage,
		Data: string(p),
	}

	for i, c := range w.conn {
		if err := c.WriteJSON(m); err != nil {
			log.Printf("ws msg write err: %v", err)

			w.conn = append(w.conn[:i], w.conn[i+1:]...)
		}
	}

	fmt.Println(m)

	return len(p), nil
}
