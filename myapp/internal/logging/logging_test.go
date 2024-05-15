package logging

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupLogging(t *testing.T) {
	logFileName := SetupLogging()

	// Check that the log file was created
	_, err := os.Stat(logFileName)
	assert.NoError(t, err)
}

func TestInitializeGinWithLogger(t *testing.T) {
	logFileName := SetupLogging()
	engine := InitializeGinWithLogger(logFileName)

	// Check that the Gin engine was created
	assert.NotNil(t, engine)
}

func TestCreateLogFile(t *testing.T) {
	logFileName := createLogFile()

	// Check that the log file was created
	_, err := os.Stat(logFileName)
	assert.NoError(t, err)
}

func TestGetLocation(t *testing.T) {
	location := getLocation("Asia/Seoul")

	// Check that the location is correct
	assert.Equal(t, "Asia/Seoul", location.String())
}

func TestCreateLogDirectory(t *testing.T) {
	createLogDirectory()

	// Check that the logs directory was created
	_, err := os.Stat("./logs")
	assert.NoError(t, err)
}
