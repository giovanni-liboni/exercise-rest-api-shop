package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {

	// Get the passoword from the terminal
	fmt.Print("Enter the password: ")
	var password string
	_, err := fmt.Scanln(&password)
	if err != nil {
		log.Fatalln(err)
	}

	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}

	// Display the encrypted password in the terminal
	fmt.Println("Encrypted password: ", string(generateFromPassword))

}
