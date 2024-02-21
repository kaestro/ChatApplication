// myapp/tests/SessionManager_test.go
package tests

import (
	session "myapp/internal/session"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionManager(t *testing.T) {
	manager := session.GetSessionManager()

	// Test SetSession
	err := manager.SetSession("testKey", "testValue")
	assert.NoError(t, err)

	// Test GetSession
	val, err := manager.GetSession("testKey")
	assert.NoError(t, err)
	assert.Equal(t, "testValue", val)

	// Test DeleteSession
	err = manager.DeleteSession("testKey")
	assert.NoError(t, err)

	// Test GetSession after deletion
	val, err = manager.GetSession("testKey")
	assert.Error(t, err)
	assert.Equal(t, "", val)
}
