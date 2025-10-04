package main

import "fmt"


func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age < 18 {
		fmt.Println("You are not an adult.")
	} else {
		fmt.Println("Invalid age entered.")
	}

	// ek achi baat aur cool ye hai ki ham variable ko declare saath hi saath uske upar condition bhi likh sakte hai! like ----
	// Aur iss declare variable ko use bhi kar sakte hai kahi bhi if me yaa else if me bhi
	if votingAge := 18; votingAge <= 20 {
		fmt.Println("You are eligible to vote.", votingAge)
	} else {
		fmt.Println("You are not eligible to vote and you are not jawaan.", votingAge)
	}
}