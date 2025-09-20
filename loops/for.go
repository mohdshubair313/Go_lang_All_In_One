package main

import "fmt"

// go ke andar for loop hi use karte hai and idhar while loop nahi hota hai
func main() {
	// pehle while loop ka example dekhte hai
	i:= 1
	for i <= 10  {
		fmt.Println(i)
		i++
	}


	// for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// break for breaking the loop and continue for continuing the loop
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break				// yaha par 5 par loop break ho jayega
		}
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue				// yaha par 5 ko hata ke 
		}
		fmt.Println(i)
	}


	// agar aapko khaali range define karni hai ki kaha tak loop chalana hai and write for loop in easy way so there is range in go

	for i:= range 5 {
		fmt.Println(i)		// --> isse 0 se start hota hai loop har baar and 4 tak print karega means 5 exclude karega
	}						// means jaha tak jaana hai bas usse ek zaada likh do
}