package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func Server(logger *log.Logger) *http.Server {

	HTTPRouter := http.NewServeMux()
	HTTPRouter.HandleFunc(`/`, handlers.HandlersRoot)
	HTTPRouter.HandleFunc(`/upload`, handlers.HandlersUpload)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      HTTPRouter,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return s

}
