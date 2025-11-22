package main
// agar multiple fields ko group karna hai to structs use kar sakte hai 
import "time"
import "fmt"

// order structs 
type order struct {
	id string
	amount float32 			// ye struct blueprint ke jaise kaam karta hai
	status string			// isse ham multiple instances bana sakte hai
	createdAt time.Time
}

func main() {	// complex data structure banana hai to structs ka use karte hai
	MyOrder:= order{
		id: "1",
		amount: 50,            // ye order ke types me values ko define karna hai
		status: "received",
	}

	MyOrder.createdAt = time.Now() // ye baad me agar values ko add karna hai to aise bahar add kar sakte hai

	// agar particular value ko get karna hai struct me se to aise get karo
	fmt.Println(MyOrder.status)


	fmt.Println("order struct", MyOrder) // pura struct ko print karwa dena

	// new instance of order struct means of order blueprint
	MyOrder2:= order {
		id: "2",
		amount: 29,
		status: "delivered",
		createdAt: time.Now(),
	}
	fmt.Println(MyOrder2)

	// so hum struct ke madad se OOPs ke tarah data ko organize kar sakte hai and multiple instances bana sakte hai jo alag alag data hold karenge

//--------***** -----------------------------------------------------------------

// structs ke ham function bhi define kar sakte hai jo particular struct ke liye kaam karega and jaise oops me behaviour define karte hai jisme methods hote hai waise 




}