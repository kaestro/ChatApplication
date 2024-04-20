package main

import (
	"flag"
	"myapp/api/routes"
	"myapp/internal/logging"

	"github.com/gin-gonic/gin"
)

func main() {
	mode := flag.String("mode", "debug", "mode in which the server should run")
	flag.Parse()

	if *mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	logFileName := logging.SetupLogging()

	r := logging.InitializeGinWithLogger(logFileName)
	routes.SetupRoutes(r)

	r.Run(":8080")
}
