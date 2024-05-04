package main

import (
	"fmt"
	"math"
	"os"

	"rsc.io/quote"
)

func main() {
	time_week()
	for_while()
	for_loop([]int{1, 5, 5, 10, 12})
}

func main_str_test() {
	doubleQuotes := "Hello \n World !"
	backQuotes := `Back \n World !`
	var singleChar int32 = 'A'

	fmt.Println(doubleQuotes)
	fmt.Println(backQuotes)
	fmt.Println(singleChar)
	fmt.Printf("%T\n", singleChar)
}

func dev() {
	fmt.Println("=======\n" + quote.Go() + "\n=======")
}

func math_test() {
	sqrt := 5
	var hint = "-Math.Sqrt-"
	fmt.Printf("%s | Number %d, Sqrt Number: %f", hint, sqrt, math.Sqrt(float64(sqrt)))
}

func input_from_console() {
	var input string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&input)
	fmt.Printf("Form console: %s\n", input)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x float64, y float64) float64 {
	return x - y
}

func env(name string) string {
	return os.Getenv(name)
}

func swap(left, right string) (string, string) {
	return right, left
}

func swap_temp_var(left, right string) (string, string) {
	temp := left // := only for initialization
	left = right
	right = temp
	return left, right
}

func print_swap() {
	fmt.Println(swap("le", "re"))
	fmt.Println(swap_temp_var("le", "re"))
}

// Named return values. This is known as a "naked" return.
func split(sum int) (x int, y int) {
	x = (sum * 4) / 9
	y = sum - x
	return
}
