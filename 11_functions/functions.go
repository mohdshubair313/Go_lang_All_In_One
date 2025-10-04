package main

import "fmt"

func twoSum(a int, b int) int { // agar bohot se parameter same type ke hi hai
	return a + b				// to ham a,b int karke bhi likh sakte hai
}

// go me ham function ko multiple return bhi kar sakte hai and agar aisa karenge to type bhi define karna padega function_name ke aage () bracket me 
// order same hona chahiye inn types ko define karne ka inn return values ke mutabik
func getlanguages() (string, string, string) {
	return "golang", "javascript", "python"
}

func getprocess(fn func(a int) int) {
	fn(1)
}

func processit() func(a int) int {
	return func(a int) int {
		return 4
	}
}



func main() {

	// fmt.Println("starting the first line of main function")

	// result:= twoSum(3,4);
	// fmt.Println(result)

	fmt.Println(getlanguages())
	// Note: aap inn teeno values ko get bhi karsakte hai like
	// lang1, lang2, lang3:= getlanguages()
	// fmt.Println(lang1, lang2, lang3)

	fn:= func(a int) int {
		return 3
	}
	fmt.Println(fn)

	fun:= processit()
	fun(6)


}