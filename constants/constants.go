package main
import "fmt"

// ek baar hamne value assign kardi to usko change nahi kar sakte (const)

func main() {
	const name string = "Hello, World!"
	fmt.Println(name)

	name = "Hello, Go!" // Error: cannot assign to name (constant variable)
	fmt.Println(name)
}