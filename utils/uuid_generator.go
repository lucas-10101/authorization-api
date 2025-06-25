package utils

import "github.com/google/uuid"

// Main method, use to allow future changes in uuid version, random or pool settings
func GenerateUUID() string {
	return uuid.NewString()
}
