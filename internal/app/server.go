package server

import (
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/handler"
)

type Server struct {
	handler *handler.Handler
}

type URL struct {
    URL string `json:"URL"`
}

func InitServer( h *handler.Handler) *Server {
	return &Server{ handler: h}
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet: 
		s.handler.HandleGetURL(w,r)

	case http.MethodPost: 
		s.handler.HandlePostURL(w,r)
		
	default: 
		http.Error(w, "Wrong request", http.StatusBadRequest)
		return
	}
}
