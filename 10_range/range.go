package main
import "fmt"

// range keyyword for iterating over data structures

func main() {
	nums := []int {12,15,18,23}
	
	sum:= 0
	for i, num := range nums {
		sum += num
		fmt.Println(i, num, sum)
	}
	println("total:", sum)


	m := map[string]int {"a":1, "b":2, "c":3}

	for k, v:= range m {
		fmt.Println(k,v)
	}
}