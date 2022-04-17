package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Router routes endpoint requests to their designated functions
var Router *mux.Router

// CreateHTTPServer applies the functions to their
func CreateHTTPServer() http.Server {
	connectionString := Config.Server.Host

	// Indicates a development setup that has a port number
	if Config.Application.Environment != "production" {
		connectionString = fmt.Sprintf("%s:%d", Config.Server.Host, Config.Server.Port)
	}

	Router.Use(LoggerMiddleware)
	return http.Server{
		Handler:      Router,
		Addr:         connectionString,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
