package main

import (
	"fmt"
	"log" // Log package for logging error messages
	"os"  // OS package used for interacting with the OS such as reading files

	"github.com/golang-jwt/jwt/v5" // Importing the JWT library for Go to handle JSON Web Tokens
)

// JWT token given as string
var tkn = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyMjU4MywiaWF0IjoxNjk3NjE5NTgzfQ.MkClwLYU-TQgzjMRCrnFoFA4ccVr6Q0zts11YTCu-k8gK9UWy1P1Rfslcvh5Q8n6wnHKJPviu6EbcaZ4ytqqf9z84wFjSK7xoUAobVzN7_1mN_lXseM_nLD4uINIVOeT0FZOafDiXkR0L39ur7JCLuZVAtfeQwcNCgmFghwKMJ9hLd9-W9SnGsIAXsppJ0ew0Bc0kVuIjsGi4P3EK-jyjsdhj82fhLYnX3ZUmQ5k8U2clolwfJLdwcGaKDTCZ40J9fJpLmJ7bHn6E9OYq7-QFzHyohdcLbJIzXYG--I19VmdsI4OEBI0eiH7J_wc9JlXAWOFXxtF0qHCemZ6bXQszw`

func main() {
	// Reads the public key from pubkey.pem file
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	var c jwt.RegisteredClaims
	// Parsing the JWT token with the claims
	token, err := jwt.ParseWithClaims(tkn, &c, func(token *jwt.Token) (interface{}, error) {
		// Provides the public key for validating the JWT token
		return publicKey, nil
	})

	if err != nil {
		// If error while parsing the token, print the error and exit
		log.Println("parsing token", err)
		return

	}
	if !token.Valid {
		// If the token is not valid, log the error and exit
		log.Println("invalid token")
		return
	}

	// Print the claims from the token
	fmt.Printf("%+v", c)

}
