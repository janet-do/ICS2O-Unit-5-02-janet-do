// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file contains the JS functions for index.html
function myButtonClicked() {
  alwaysOnButtonChecked = document.getElementById("on-check").checked

  if (alwaysOnButtonChecked == true) {
    document.getElementById("radio-button-value").innerHTML =
      "<p>Value is: On</p>"
  } else {
    document.getElementById("radio-button-value").innerHTML =
      "<p>Value is: Off</p>"
  }
}
