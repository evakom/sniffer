package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/evakom/sniffer/internal/models"
	"github.com/gorilla/websocket"
)

const pingTimeOut = 5 * time.Second

// ApplyHandlers applies all handlers.
func (srv *Server) ApplyHandlers() {
	srv.router.Handle("/*", http.FileServer(http.Dir("./web")))
	srv.router.Get("/socket", srv.socketHandler)
}

func (srv *Server) socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}

	srv.Writer.conn = append(srv.Writer.conn, conn)

	go func() {
		for {
			<-time.After(pingTimeOut)

			msg := models.Message{Type: models.MTPing}
			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
				break
			}
		}
	}()

	for {
		msg := models.Message{}
		if err := conn.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Fatalf("ws msg read err: %v", err)
			}

			break
		}

		if msg.Type == models.MTPong {
			continue
		}
	}

	fmt.Println("WebSocket closed")
}
