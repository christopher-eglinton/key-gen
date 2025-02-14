package main

import (
	"testing"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TestCreateNewKey checks if createNewKey generates a valid UUID
func TestCreateNewKey(t *testing.T) {
	key := createNewKey()

	// Ensure the UUID is valid
	_, err := uuid.Parse(key)
	if err != nil {
		t.Errorf("Generated key is not a valid UUID: %s", key)
	}
}

// TestHashKey checks if the hashKey function works correctly
func TestHashKey(t *testing.T) {
	apiKey := "supersecret123"

	// generate the hashed key
	hashedKey, err := hashKey(apiKey)
	if err != nil {
		t.Fatalf("Error hashing key: %v", err)
	}

	// verify hash matches the original key
	err = bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(apiKey))
	if err != nil {
		t.Errorf("Hash did not match the original key: %v", err)
	}

	// verify that a different key does not match the hash
	wrongKey := "wrongpassword"
	err = bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(wrongKey))
	if err == nil {
		t.Errorf("Hash matched an incorrect password, which should not happen")
	}
}

// BenchmarkHashKey tests the performance of hashKey function
func BenchmarkHashKey(b *testing.B) {
	apiKey := "supersecret123"
	for i := 0; i < b.N; i++ {
		_, _ = hashKey(apiKey)
	}
}
