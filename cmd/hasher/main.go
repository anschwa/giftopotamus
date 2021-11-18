package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./hasher PASSWORD")
		os.Exit(1)
	}

	h, err := hash(os.Args[1])
	if err != nil {
		fmt.Println("Unable to generate password hash:", err)
		os.Exit(1)
	}

	fmt.Println(h)
}

// hash takes a plaintext password and return a bcrypt hash
func hash(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(h), nil
}
