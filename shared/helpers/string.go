package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomString(length int) (string, error) {
	// Create a byte slice to hold the random data
	bytes := make([]byte, length)
	// Read random bytes from the crypto/rand package
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random string: %w", err)
	}
	// Encode the random bytes to a base64 string
	return base64.URLEncoding.EncodeToString(bytes), nil
}
