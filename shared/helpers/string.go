package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	// Calculate the number of random bytes needed
	// Base64 encoding: 4 characters represent 3 bytes
	// So we need (desired_length * 3) / 4 bytes
	randomBytes := make([]byte, (length*3)/4)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}

	encoded := base64.URLEncoding.EncodeToString(randomBytes)
	// Trim to exact length requested
	return encoded[:length], nil
}
