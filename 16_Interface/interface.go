package main

import "fmt"

// ==========================================
// 1. INTERFACE DEFINITION (Contract)
// ==========================================

// 'paymenter' ek interface hai. Ye ek 'Contract' (shart) ki tarah hota hai.
// Iska kehna hai: "Jo bhi struct 'paymenter' banna chahta hai, 
// uske paas 'pay(amount float32)' naam ka function hona ZAROORI hai."
type paymenter interface {
	pay(amount float32)
}

// ==========================================
// 2. DEPENDENCY INJECTION STRUCT
// ==========================================

// 'payment' struct humara main logic handler hai.
// Note: Iske andar hum 'razorpay' ya 'stripe' ka naam nahi likh rahe.
// Hum yahan 'paymenter' interface likh rahe hain.
// Isse hume faida ye hoga ki future mein hum Razorpay ko hata kar Stripe bhi laga sakenge bina code tode.
type payment struct {
	gateway paymenter // Yahan wo object store hoga jo paymenter interface ko implement karta hai.
}

// Ye 'payment' struct ka method hai.
func (p payment) makePayment(amount float32) {
	// Puran tareeka (Bad Practice - Tight Coupling):
	// razorpaymentGw := razorpay{}
	// razorpaymentGw.pay(amount)
	// ^ Agar aisa karte, toh kal ko Stripe lagane ke liye ye pura function badalna padta.

	// Sahi tareeka (Good Practice - Loose Coupling):
	// Hum 'p.gateway' use kar rahe hain. Hume nahi pata ye Razorpay hai ya kuch aur.
	// Hume bas itna pata hai ki iske paas 'pay' function hai kyunki ye ek interface hai.
	p.gateway.pay(amount)
}

// ==========================================
// 3. CONCRETE IMPLEMENTATION (Razorpay)
// ==========================================

// 'razorpay' ek struct hai jo actual kaam karega.
type razorpay struct{}

// Yahan 'razorpay' struct 'pay' method ko implement kar raha hai.
// Go language automatically samajh leti hai ki kyunki 'razorpay' ne 
// 'pay' method define kar diya hai, isliye ab 'razorpay' ek 'paymenter' bhi ban gaya hai.
func (r razorpay) pay(amount float32) {
	// Actual logic yahan likha jata hai (API calls, etc.)
	fmt.Println("Making payment using Razorpay:", amount)
}

// (Extra Example): Agar kal ko Stripe add karna ho toh bas naya struct banao:
// type stripe struct {}
// func (s stripe) pay(amount float32) { fmt.Println("Stripe se payment hui", amount) }

// ==========================================
// 4. MAIN FUNCTION (Execution)
// ==========================================

func main() {
	// Step 1: Hum 'razorpay' ka object bana rahe hain.
	// Ye concrete implementation hai.
	myRazorpayInstance := razorpay{}

	// Step 2: Hum 'payment' struct bana rahe hain.
	// Yahan hum 'razorpay' ke object ko 'paymenter' interface ki jagah pass kar rahe hain.
	// Ye valid hai kyunki razorpay ke paas 'pay()' function hai.
	newPayment := payment{
		gateway: myRazorpayInstance,
	}

	// Step 3: Ab hum payment process kar rahe hain.
	// 'makePayment' function ko ye nahi pata ki andar Razorpay hai, wo bas apna kaam kar raha hai.
	newPayment.makePayment(100)
}