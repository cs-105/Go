package main

import "fmt" 
  

// main function
func main() {

	fmt.Println("Welcome to Interactive Madlibs!")
  
    // Println function is used to
    // display output in the next line
    fmt.Println("To begn, please enter a topic: ")
  
    // var then variable name then variable type
    var userInput string
  
    // Taking input from user
    fmt.Scanln(&userInput)
  
    // Print function is used to
    // display output in the same line
    fmt.Print("You have selected: ")
  
    // Addition of two string
    fmt.Print(userInput)
}