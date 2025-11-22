package main
import "fmt"

func main() {
	var name string = "John"  //  var is the keyword to declare a variable, name is the variable name, string is the type of variable and "John" is the value assigned to the variable
	fmt.Println(name)

	// agar ham variable ko string type naa de tab bhi go ko pata chal jaata hai ki ye string hai means infer kar leta hai like
	var city = "New York"
	var isAdult = true
	var age = 30 
	fmt.Println(age)
	fmt.Println(isAdult)
	fmt.Println(city)

	// ham shorthand ka use bhi kar sakte hain use karte hai jab variable declare bhi karna hai and type bhi declare karna hai
	country := "USA"
	fmt.Println(country)

	// phir aise to kabhi type declare karna hi nahi padega ... par aisa nahi hai kuch jagah par ham khaali variable declare karte hai and value defined nahi karte to usme type define karna padega like this
	var state string

	state = "California"
	fmt.Println(state)

}