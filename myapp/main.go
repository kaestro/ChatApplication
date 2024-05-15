package main

import (
	"flag"
	"log"
	"myapp/api/routes"
	"myapp/internal/logging"
	"myapp/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	mode := flag.String("mode", "debug", "mode in which the server should run")
	flag.Parse()

	if *mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	logFileName := logging.SetupLogging()
	logFile, _ := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Server is starting...")

	r := logging.InitializeGinWithLogger(logFileName)
	routes.SetupRoutes(r)
	middleware.SetupMiddleware(r)

	r.Run(":8080")

	log.Println("Server has stopped.")
}
