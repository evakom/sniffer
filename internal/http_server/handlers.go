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
	ws, err := srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}

	// ------------------
	//srv.Writer, err = ws.NextWriter(websocket.TextMessage)
	//if err != nil {
	//	log.Printf("error get next websocket writer: %v", err)
	//}

	//m := models.Message{
	//	Type: models.MTMessage,
	//	Data: "qqq www eee",
	//}

	//if err := ws.WriteJSON(m); err != nil {
	//	log.Printf("ws msg write err: %v", err)
	//}

	//_, err = srv.Writer.Write([]byte(`{"Type":"message", "Data":"aaa sss ddd"}`))
	//if err != nil {
	//	log.Printf("error write to websocket: %v", err)
	//}

	err = ws.WriteMessage(websocket.TextMessage, []byte(`{"type": "message", "data": "aaa sss ddd"}`))
	if err != nil {
		log.Printf("error write to websocket: %v", err)
	}
	// ------------------

	go func() {
		for {
			<-time.After(pingTimeOut)

			msg := models.Message{Type: models.MTPing}
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
				break
			}
		}
	}()

	//serv.subscribers[id] = func(msg string) error {
	//
	//	return nil
	//}

	for {
		msg := models.Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Fatalf("ws msg read err: %v", err)
			}

			break
		}

		if msg.Type == models.MTPong {
			fmt.Println("browser pong received")
			continue
		}
	}

	fmt.Println("WebSocket closed")
}
