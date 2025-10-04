package main

import (
	"fmt"
	"maps"
)

// python me dictionary jaise hoti hai waise hi hote hai maps
// javascript me hash, object wagera hote hai wo maps in go hai
// Note:  the make() function is used for initializing slices, maps, and channels
func main() {
	// why we not initialize map like slice like var m =  map[string]int
	//When we declare a map as nil (var m map[KeyType]ValueType), it is uninitialized and not yet ready to store key-value pairs. So this will create a runtime error when we try to add values to a nil map. To avoid such errors create map using make
	m := make(map[string]string)

	// setting elements
	m["name"] = "holand"
	m["age"] = "24"

	//---------> get an element

	fmt.Println(m["name"], m["age"])  // print me key daalo like name and uska value hame mil jayega like dict in python

	fmt.Println(m["phone"]) // phone key hai hi nahi so zero value aayega like string me empty string "", integrer me 0 and bool me false

	//---------> deleting an element
	delete(m, "age")
	fmt.Println(m)

	// --------------> technique to check if a key is present in map or not
	Map := map[string]int{"price": 100, "age": 24}

	value, ok :=Map["price"]  // ok me true aayega agar key present hai to otherwise false
	fmt.Println(value, ok)

	// value, ok =Map["phone"]  // phone key hai hi nahi to ok me false aayega
	// fmt.Println(value, ok)

	if ok {
		fmt.Println("key is present")
	}else {
		fmt.Println("key is not present")
	}

	// ---------> map standard package 
	m1 := map[string]int {"price": 100, "old": 3, "phone": 200}
	m2 := map[string]int {"price": 100000, "old": 4, "phone": 150}
	
	fmt.Println(maps.Equal(m1,m2))
}
