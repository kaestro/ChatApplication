// /myapp/internal/logging/logging.go
package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	locationSeoul   = "Asia/Seoul"
	locationDefault = "UTC"
)

func SetupLogging() string {
	createLogDirectory()
	return createLogFile()
}

func InitializeGinWithLogger(logFileName string) *gin.Engine {
	f, _ := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	return r
}

func createLogFile() string {
	loc := getLocation(locationSeoul)
	currentTime := time.Now().In(loc)
	logFileName := fmt.Sprintf("./logs/gin_%s.log", currentTime.Format("2006_01_02"))
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		os.Create(logFileName)
	}
	return logFileName
}

// getLocation returns the time.Location for the given location string.
// If the location is invalid, it returns UTC instead.
func getLocation(location string) *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("Failed to load location %s: %v. Trying UTC instead.", location, err)
		loc, err = time.LoadLocation(locationDefault)
		if err != nil {
			log.Fatalf("Failed to load location UTC: %v", err)
		}
	}
	return loc
}
func createLogDirectory() {
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", os.ModePerm)
	}
}
