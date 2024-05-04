package main

import "fmt"

func spread(arr ...int) {
	fmt.Println("Spread example")
	for i := 0; i < len(arr); i = i + 1 {
		fmt.Printf("Index: %d | Value %d\n", i, arr[i])
	}
	fmt.Println("===============")
}

func spread_invoke() {
	spread(2, 3, 4, 10)

	array := []int{99, 99, 99}

	spread(array...)

	array_d := make([]int, 0, 3) // (type, start length, capacity)
	array_d = append(array_d, 9, 8, 7)

	fmt.Println(array_d)

	spread(array_d...)
}
