package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	key := flag.String("key", "empty", "key you want to encrypt.")
	flag.Parse()

	if *key == "empty" {
		fmt.Println("please enter a key.")
		return
	}

	hashedKey, err := hashKey(*key)
	if err != nil {
		fmt.Println("error hashing your key: ", err)
	}

	fmt.Println("please save your hashed key: ", hashedKey)
}

func hashKey(apiKey string) (string, error) {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(apiKey), 14)

	return string(hashedKey), err
}
