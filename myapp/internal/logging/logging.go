// /myapp/internal/logging/logging.go
package logging

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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
	loc, _ := time.LoadLocation("Asia/Seoul")
	currentTime := time.Now().In(loc)
	logFileName := fmt.Sprintf("./logs/gin_%s.log", currentTime.Format("2006_01_02"))
	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		os.Create(logFileName)
	}
	return logFileName
}

func createLogDirectory() {
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", os.ModePerm)
	}
}
