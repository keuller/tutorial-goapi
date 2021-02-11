package web

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Server structure to represent HTTP server instance
type Server struct {
	engine *http.Server
}

// NewServer - returns an instance of http server
func NewServer(port string) *Server {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5, "application/json"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{
			"message": "It works!",
		})
	})

	router.Route("/v1/tasks", func(r chi.Router) {
		r.Get("/", listTaskHandler)
		r.Get("/{id}", readByIDHandler)
		r.Post("/", createHandler)
		r.Put("/{id}", updateHandler)
		r.Put("/{id}/toggle", toggleTaskHandler)
		r.Delete("/{id}", removeHandler)
	})

	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &Server{
		engine: server,
	}
}

// Start invoke server listening for requests
func (s *Server) Start() error {
	if err := s.engine.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.engine.Handler.ServeHTTP(w, r)
}
