package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := []byte("admin")

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))
}
