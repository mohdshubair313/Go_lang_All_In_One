package main
import "fmt"

func sum(nums ...int) int { // yaha ...int matlab hai ki ye int type ke n paramters recieve karega
	total := 0
	for _, num := range nums {  // _, ye cheez index ya value ko ignore karne ke liye hota hai
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3)) 
	// variadic function matlab hai ki ham functions n number parameters pass karsakte hai koi bhi type

	// ham values ko slice ke andar se bhi pass kar sakte hai like --
	numbers := []int{1, 2, 3, 4, 5}
	result := sum(numbers...) // yaha ... ka matlab hai ki ye slice ke andar ke elements ko alag alag parameters ke roop me pass kar dega
	fmt.Println(result)


}