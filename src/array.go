package main

import "fmt"

type Block struct {
	size int
}

func array_fixed() {
	var fix []string = []string{"vlad", "hello"}

	fmt.Println(fix)

	fix = append(fix, "appendedValue")

	fmt.Println(fix)

	// Empty initializer
	fixed_empty := []string{}
	fmt.Println(fixed_empty)
}

func array_make() {
	// dynamic sized array.
	maked_array := make([]string, 2, 10)
	fmt.Println("Len: ", len(maked_array))
	maked_array[0] = "vlad"
	maked_array[1] = "other"
	fmt.Println("Len: ", len(maked_array))
}
