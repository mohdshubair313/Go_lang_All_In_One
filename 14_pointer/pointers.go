package main
import "fmt"

func changeNum(num int) {
	num = 5;
	println(num)
}							// this is the example of pass by value where changeNum func me jo value hai num wo "change" nahi kar sakta func main me bhi 
								// because dono me variable alag jagah par store ho rahe hai

func changeNumWithPointer(num *int) int { // "*" jo hai num variable jo recieve hoga uska address hai basically
	*num = 1 // "*num" bata raha hai ki jo num variable hai uske address par jaake value change karna hai
	fmt.Println("The change number is this: ",*num) // isme star laga do because dereference ho jayega and hame value dhikh jayegi
	return *num

}

func main() {
	// pass by value example (Isme value change nahi hoti hai)
	num:= 3
	changeNum(num)

	fmt.Println("the number before is ", num)

	// pass by reference example
	changenum := changeNumWithPointer(&num) // "&num" because ye function ek pointer expect kar rahi hai means pointer and memory address "&" lagake get karte hai 
	fmt.Println("the number after is ", changenum)	

}

