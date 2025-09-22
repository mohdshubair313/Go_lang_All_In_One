package main

import (
	"fmt"
	"slices"
)

func main() {
	// slices are the dynamically array and most used in go
	// we also can use useful method like to add element or remove like that

	/*
	var nums []int
	fmt.Println(nums == nil) // returns true because nums arr me kuch nahi hai

	var nums = make([]int, 2)

	fmt.Println(nums)
	*/

//--- --- --------------------------------------------------------------------------------------------------------------------

	/*
	var nums = make([]int, 2,  5) // here 2 is the size means length kitni hogi slice ki initially and 5 is the capacity 

	fmt.Println(cap(nums)) // returns the maximum numbers of elements can fit but it doesn't means that it only size of z but it can rezise dynamically

	nums = append(nums, 1) // here we are adding the element in the slice
	nums = append(nums, 2)
	fmt.Println(nums)
	// jab print hoga slice to initially 0 0 because length hi 2 se start kari hai so isi liye but agar 0 se start kar denge to jo append karenge wo hi dhikhayega 
	*/

	//--- --- ----------------------------------------------------------------------------------------------------------------------------------
	
	/*
	// ham slice ko copy bhi kar sakte hai like
	var slice = make([]int, 0, 5)
	slice = append(slice, 1)
	slice = append(slice, 4)
	
	fmt.Println(slice)
	var NewSlice = make([]int, len(slice))
	// copy(NewSlice, slice)
	fmt.Println(NewSlice)
	*/

	//--- --- ----------------------------------------------------------------------------------------------------------------------------------

	// slice operator for agar hame kuch specific elements get karne hai to slice operator use karte hai
	/*
	var Nums = []int{1,2,3,4,5,6}

	fmt.Println(Nums[0:2]) // this Nums[1:2] means from - to matlab 1 se 2 index tak jana par 2 index ko chod do
	fmt.Println(Nums[1:]) // this means from 1 index to last index
	fmt.Println(Nums[:4]) // this means from 0 index to 4 index but 4 index ko chod do
	*/

	//--- --- ----------------------------------------------------------------------------------------------------------------------------------

	// slices standard library package built in go
	//har slice like arr1 and arr2 ke elements ko check karta hai ki same hai ki nahi agar saare same hote hai to true otherwise false 
	/*
	var arr1  = []string{"shaikh", "shubair"}
	var arr2  = []string{"shaikh", "shubair"}

	fmt.Println(slices.Equal(arr1, arr2))
	*/
	//--- --- ----------------------------------------------------------------------------------------------------------------------------------	

	// 2D slices 
	var nums= [][]int{{1,2}, {3,4}}
	fmt.Println(nums)
}