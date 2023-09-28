package server

import (
	"net/http"

	"github.com/devnandito/servergenpwd/models"
)

type Server struct {
	port string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...models.Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) File(dir string) {
	style := http.FileServer(http.Dir(dir))
	http.Handle("/assets/", http.StripPrefix("/assets/", style))
}