package main

import (
	"log"

	"github.com/mercuryheart123/lmwn-backend-go/server"
	"go.uber.org/zap"
)

func main() {
	var port string = "8080"
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("Failed to initialize zap logger")
	}
	defer logger.Sync()

	app := server.InitServer(port, logger)
	app.InitRoutes()
	app.RunServer()

}
