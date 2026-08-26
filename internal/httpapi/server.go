package httpapi

import (
	"net/http"
	"spectralcalibrationlab/internal/lab"
)

type Server struct {
	engine *lab.Engine
	mux    *http.ServeMux
}

func New(engine *lab.Engine) *Server {
	server := &Server{engine: engine, mux: http.NewServeMux()}
	server.routes()
	return server
}

func (s *Server) Handler() http.Handler { return s.mux }

func (s *Server) routes() {
	s.mux.HandleFunc("GET /healthz", s.health)
	s.mux.HandleFunc("GET /v1/profiles", s.profiles)
	s.mux.HandleFunc("POST /v1/observations", s.observations)
	s.mux.HandleFunc("GET /v1/reports", s.reports)
}
