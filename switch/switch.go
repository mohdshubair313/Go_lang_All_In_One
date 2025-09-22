package main
import "fmt"

func main() {
	// switch tab lagao jab bohot saare conditions ho
	// break lagane ki jarurat nahi hoti wo go apne aap handle kar leta hai
	// agar multiple conditions same hai to unhe comma se alag kar do

	fmt.Println("Enter a subject name")
	var i string
	fmt.Scan(&i)

	switch i {
	case "English":
		fmt.Println("English is bhery bhery eajjy subject")
	case "Maths", "maths":
		fmt.Println("Maths is also eajjy subject")
	case "Science", "science":
		fmt.Println("Science is also eajjy subject")
	default:
		fmt.Println("Unknown subject is wrong ..")
	}


	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am int")
		case string:
			fmt.Println("I am string")
		case bool:
			fmt.Println("I am bool")
		default:
			fmt.Println("Unknown type")
		}
	}
	whoAmI(7)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)
}