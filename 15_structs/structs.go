package main

// agar multiple fields ko group karna hai to structs use kar sakte hai
import (
	"fmt"
	"time"
)

// order structs
type order struct {
	id string
	amount float32 			// ye struct blueprint ke jaise kaam karta hai
	status string			// isse ham multiple instances bana sakte hai
	createdAt time.Time
}

func main() {	// complex data structure banana hai to structs ka use karte hai

	myOrder := newOrder("123", 250.50, "processing") // ye constructor function ka use karke bana hai
	fmt.Println(myOrder)

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

	MyOrder2.changestatus("Done dana dan")
	fmt.Println("after changing status", MyOrder2)


	// agar aapko ek hi baar struct banake use karna hai so aap inline structs bana sakte ho

	language := struct {
		name string
		difficultyLevel string
	}{"hindi", "medium"}
	fmt.Println("inline struct", language)

}

// structs ke ham function bhi define kar sakte hai jo particular struct ke liye kaam karega and jaise oops me behaviour hote hai waise hi go me bhi hum struct ke liye functions define kar sakte hai

/*

func (o order) changestatus(status string) {		// ye order struct ke liye function hai jo status change karega and ye jo func ke aage hai wo reciever hai jo batata hai ki ye function kis struct ke liye hai
	o.status = status	// ab isme status change nahi hoga q? because ye ek copy hai original struct ka, agar original struct ko change karna hai to pointer use karna padega
}

*/


// so ab isko pointer banane ke liye o ke aage * laga do to make this as pointer
func (o *order) changestatus(status string) {
	o.status = status	// and deferencing ka kaam ye khud karta hai so no need to put & before variable
}

func (o order) getAmount() float32 {
	return o.amount
}

// so * tabhi use karo jab struct ki kisi value ko modify karna ho par agar modify nahi karna to second tareke se jaise likha hai usme star (*) use karni ki zoroorat nahi

// ----------------- *****************************----------------------------------

// Contructor kaise bana sakte hai Go lang me

func newOrder(id string, amount float32, status string) *order {

	myOrder := order {
		id:id,
		amount: amount,
		status: status,
	}

	return &myOrder		// ye pointer return karega
}

// agar aapko structs ke andar ek aur struct daalna hai so aap ye kar sakte hai bas ek struct ke andar dusre struct ka naam daal ke
type customer struct {
	name string
	phone string
}

type custOrder struct {
	id string
	address string
	customer
}