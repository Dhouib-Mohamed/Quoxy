package id

import "github.com/google/uuid"

func GenerateRandomId() string {
	// Generate a random id
	newUUID := uuid.New().String()
	if len(newUUID) > 16 {
		newUUID = newUUID[:16]
	}
	return newUUID
}
