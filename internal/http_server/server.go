package httpserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

const wsBufferSize = 1024

// Server base struct.
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader
	conn     *websocket.Conn
	Writer   *writer
}

// New returns new server.
func New() *Server {
	router := chi.NewRouter()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  wsBufferSize,
		WriteBufferSize: wsBufferSize,
	}

	serv := &Server{
		router:   router,
		upgrader: upgrader,
		Writer: &writer{
			conn: new(websocket.Conn),
		},
	}

	serv.ApplyHandlers()

	return serv
}

// Start starts server.
func (srv *Server) Start() error {
	return http.ListenAndServe(":8080", srv.router)
}
