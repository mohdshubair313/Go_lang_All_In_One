// interface ek tarah ke contract hote hai


package main
import "fmt"

type payment struct {		// ye payment struct hai jo data structure define karega
}

func (p payment) makePayment(amount float32) {
	razorpaymentGw:= razorpay{}
	razorpaymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

func main() {
	payment:= payment{}
	payment.makePayment(100)
}