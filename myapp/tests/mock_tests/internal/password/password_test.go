package password

import (
	"myapp/internal/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {

	t.Run("should hash a password without error", func(t *testing.T) {
		pass := "mySecretPassword"
		hashed, err := password.HashPassword(pass)

		assert.NoError(t, err)
		assert.NotEqual(t, pass, hashed)
	})
}

func TestCheckPasswordHash(t *testing.T) {

	t.Run("should return true for matching password and hash", func(t *testing.T) {
		pass := "mySecretPassword"
		hashed, err := password.HashPassword(pass)
		if err != nil {
			t.Fatalf("Hashing password failed: %v", err)
		}

		match := password.CheckPasswordHash(pass, hashed)

		assert.True(t, match)
	})

	t.Run("should return false for non-matching password and hash", func(t *testing.T) {
		pass := "mySecretPassword"
		wrongPass := "myWrongPassword"
		hashed, err := password.HashPassword(pass)
		if err != nil {
			t.Fatalf("Hashing password failed: %v", err)
		}

		match := password.CheckPasswordHash(wrongPass, hashed)

		assert.False(t, match)
	})
}
