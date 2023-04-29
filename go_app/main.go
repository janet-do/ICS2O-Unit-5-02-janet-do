// Created by: Janet
// Created on: Sep 2020
//
// The positive negative number selector
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var numberType string

	// Prompt user to select number type
	fmt.Println("Select number type: 'positive' or 'negative'")
	fmt.Scanln(&numberType)

	var randomNumber int64
	switch numberType {
	case "positive":
		// Generate a random positive number
		n, _ := rand.Int(rand.Reader, big.NewInt(6))
		randomNumber = n.Int64() + 1
	case "negative":
		// Generate a random negative number
		n, _ := rand.Int(rand.Reader, big.NewInt(6))
		randomNumber = -1 * (n.Int64() + 1)
	default:
		fmt.Println("Invalid input. Please enter 'positive' or 'negative'.")
		return
	}

	fmt.Printf("Random number is: %d\n", randomNumber)
}
