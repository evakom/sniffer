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
	var err error

	srv.conn, err = srv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}

	srv.Writer.conn = srv.conn
	srv.Writer.upgraded = true

	// ------------------
	//srv.Writer, err = ws.NextWriter(websocket.TextMessage)
	//if err != nil {
	//	log.Printf("error get next websocket writer: %v", err)
	//}

	//m := models.Message{
	//	Type: models.MTMessage,
	//	Data: string(srv.Writer.data),
	//}
	//
	//if err := srv.conn.WriteJSON(m); err != nil {
	//	log.Printf("ws msg write err: %v", err)
	//}
	//
	//fmt.Println(m)
	//srv.Writer.Close()

	//_, err = srv.Writer.Write([]byte(`{"Type":"message", "Data":"aaa sss ddd"}`))
	//if err != nil {
	//	log.Printf("error write to websocket: %v", err)
	//}

	//err = ws.WriteMessage(websocket.TextMessage, []byte(`{"type": "message", "data": "aaa sss ddd"}`))
	//if err != nil {
	//	log.Printf("error write to websocket: %v", err)
	//}
	// ------------------

	go func() {
		for {
			<-time.After(pingTimeOut)

			msg := models.Message{Type: models.MTPing}
			if err := srv.conn.WriteJSON(msg); err != nil {
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
		if err := srv.conn.ReadJSON(&msg); err != nil {
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
