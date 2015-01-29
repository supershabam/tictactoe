package tictactoe

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Addr   string
	Public string
}

func (s Server) Run() error {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(s.Public))))
	mux.HandleFunc("/statestream.ws", func(w http.ResponseWriter, r *http.Request) {
	})
	return http.ListenAndServe(s.Addr, mux)
}
