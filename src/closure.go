package main

import "fmt"

func closure_ex() {
	var i int = 0

	f := func() func() int {
		i = 10
		return func() int {
			i = 20
			return 10
		}
	}
	fmt.Printf("Before call f() i = %d\n", i)
	ff := f()

	fmt.Printf("Intermediate call f() i = %d\n", i)

	ff()
	fmt.Printf("After call ff() i = %d\n", i)
}
