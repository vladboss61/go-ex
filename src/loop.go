package main

import "fmt"

func for_loop(arr []int) {
	for i := 0; i <= len(arr); i++ {
		fmt.Printf("Index: %d", i)
		fmt.Println()
	}
}

func for_while() {
	i := 0
	for i < 10 {
		i = i + 2
		fmt.Println(i)
	}
}

func for_single_arg() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func for_range() {
	arr := []string{"99", "99", "99"}

	for i, v := range arr {
		fmt.Printf("Index: %d | Value: %s", i, v)
	}
}
