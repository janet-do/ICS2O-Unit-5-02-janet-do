// Created by: Janet Do
// Created on: Sep 2020
// This file contains the JS functions for index.html

function myButtonClicked() {
  const positiveButtonChecked =
    document.getElementById("positive-check").checked

  let randomNumber
  if (positiveButtonChecked) {
    // Generate a random positive number
    randomNumber = Math.floor(Math.random() * 6) + 1
  } else {
    // Generate a random negative number
    randomNumber = -1 * (Math.floor(Math.random() * 6) + 1)
  }

  document.getElementById("random-number-value").innerHTML =
    "<p>Random Number is: " + randomNumber + "</p>"
}
