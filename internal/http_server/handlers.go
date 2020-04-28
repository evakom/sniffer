package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/evakom/sniffer/internal/models"
)

const pingTimeOut = 5 * time.Second

// ApplyHandlers applies all handlers.
func (serv *Server) ApplyHandlers() {
	serv.router.Handle("/*", http.FileServer(http.Dir("./web")))
	serv.router.Get("/socket", serv.socketHandler)
}

func (serv *Server) socketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}

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
	//	m := models.Message{
	//		Type: models.MTMessage,
	//		Data: msg,
	//	}
	//
	//	if err := ws.WriteJSON(m); err != nil {
	//		log.Printf("ws msg fetch err: %v", err)
	//	}
	//
	//	return nil
	//}

	//for {
	//	msg := models.Message{}
	//	if err := ws.ReadJSON(&msg); err != nil {
	//		if !websocket.IsCloseError(err, 1001) {
	//			log.Fatalf("ws msg read err: %v", err)
	//		}
	//		break
	//	}
	//
	//	if msg.Type == models.MTPong {
	//		continue
	//	}
	//
	//	if msg.Type == models.MTMessage {
	//		fmt.Println(msg.Data)
	//		serv.Lock()
	//		for _, sub := range serv.subscribers {
	//			if err := sub(msg.Data); err != nil {
	//				log.Fatalf("ws msg subs err: %v", err)
	//			}
	//		}
	//		serv.Unlock()
	//	}
	//}

	fmt.Println("WebSocket closed")
}
