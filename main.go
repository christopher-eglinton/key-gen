package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	key := flag.String("key", "empty", "key you want to encrypt.")
	newKey := flag.Bool("new-key", false, "would you like a new key?")
	flag.Parse()

	// generate new key if requested
	if *newKey {
		generatedKey := createNewKey()
		fmt.Println("Generated Key:", generatedKey)
		*key = generatedKey
	} else if *key == "" {
		fmt.Println("Error: Please provide a key or use the --new-key flag.")
		return
	}

	hashedKey, err := hashKey(*key)
	if err != nil {
		fmt.Println("error hashing your key: ", err)
	}

	// print both the key and its hash
	fmt.Println("Key:", *key)
	fmt.Println("Hashed Key:", hashedKey)
}

func createNewKey() string {
	return uuid.New().String()
}

func hashKey(apiKey string) (string, error) {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(apiKey), 14)

	return string(hashedKey), err
}
