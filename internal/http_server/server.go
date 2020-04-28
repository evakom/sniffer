package httpserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

const wsBufferSize = 1024

// Server base struct.
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader
	Writer   *writeCloser
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
		Writer:   &writeCloser{},
	}

	serv.ApplyHandlers()

	return serv
}

// Start starts server.
func (srv *Server) Start() error {
	return http.ListenAndServe(":8080", srv.router)
}

type writeCloser struct {
	io.WriteCloser
	data []byte
}

func (w *writeCloser) Write(p []byte) (int, error) {
	w.data = append(w.data, p...)

	fmt.Println(string(w.data))

	return len(p), nil
}

func (w *writeCloser) Close() error {
	w.data = w.data[:0]

	return nil
}
