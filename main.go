package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/jeanguel/street-critters/api/config"
	"github.com/jeanguel/street-critters/api/controllers"
)

func main() {
	config.InitializeApplication()

	controllers.ApplyRoutes()
	server := config.CreateHTTPServer()

	config.MainLogger.Info.Println("Starting server")
	defer func() {
		config.MainLogger.Info.Println("Server shutting down")

		server.Shutdown(context.TODO())
		config.CloseApplication()

		os.Exit(0)
	}()

	go func() {
		// FIXME: Is there another way to check if the error is simply the application terminating?
		if err := server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			config.MainLogger.Error.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
