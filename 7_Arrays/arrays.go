package main

import "fmt"

func main() {
	var nums [5]int
	// array store same type of data
	// matlab var array_name [size] type
	fmt.Println(nums)
	// agar array me kuch nahi daaloge to by default go 0 daal dega according the size kitna hai like 5 tha size to 5 zeroes daal dega and maan lo type hai boolean to false daal dega
	fmt.Println(len(nums))

	var value [4]string
	fmt.Println(value) // by default empty string daal dega

	// add kaise karenge array me values simply
	nums[0] = 10
	// agar koi element ko get karna hai to
	fmt.Println(nums[0])

	// jab ham array ko banate ho tabhi values ko add karna hai so wo karenge aise
	arr:= [2]int{1,2}
	fmt.Println(arr)

	// 2D arrays
	arr2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d)

	// Array --> fixed size pata ho to use array because we can optimize memory
	// but jab agar slices me memory dynamically allocate karta hai go
}