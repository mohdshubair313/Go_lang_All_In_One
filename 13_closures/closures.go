package main

import (
	"fmt"
)

func counter() func() int {  // yaha par function counter hai jo recieve nahi kar rahi par ek func return kar rahi hai and ye ek integer return karega
	count := 0

	return func() int {		// yaha par outer fn me jo count variable hai wo outer hai yaani bahar se ayaa hai so iska execution chalta rahega
		count  += 1
		return count
	}
}

func main() {
	
	increment := counter()	// jab koi fn execute hota hai to uske variables callstack par jaate hai and phir sab khatam hojata hai so phir se execute nahi kar sakte par yaha par aisa nahi hoga

	fmt.Println(increment())
	fmt.Println(increment())
	
}