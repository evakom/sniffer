package httpserver

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/evakom/sniffer/internal/models"
	"github.com/gorilla/websocket"
)

type writer struct {
	io.Writer
	conn     *websocket.Conn
	upgraded bool
}

func (w *writer) Write(p []byte) (int, error) {
	m := models.Message{
		Type: models.MTMessage,
		Data: string(p),
	}

	if !w.upgraded {
		return 0, errors.New("websocket connection is not open")
	}

	if err := w.conn.WriteJSON(m); err != nil {
		log.Printf("ws msg write err: %v", err)
	}

	fmt.Println(m)

	return len(p), nil
}
