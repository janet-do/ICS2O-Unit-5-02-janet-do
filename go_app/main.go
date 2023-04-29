// Created by: Janet
// Created on: Sep 2020
//
// The positive negative number selector
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numberType string
	fmt.Print("Enter 'positive' or 'negative': ")
	fmt.Scan(&numberType)

	var randomNumber int
	if numberType == "positive" {
		// Generate a random positive number
		randomNumber = rand.Intn(6) + 1
	} else if numberType == "negative" {
		// Generate a random negative number
		randomNumber = (rand.Intn(6) + 1) * -1
	} else {
		fmt.Println("Invalid number type entered.")
		return
	}

	fmt.Printf("Random Number is: %d\n", randomNumber)
}
