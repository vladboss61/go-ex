package main

import "fmt"

func pointer_basic() {
	value := 25
	var ptr_value *int = &value
	fmt.Printf("Before Changes, Ptr Address: %d |  Ptr Value: %d | Value: %d\n", ptr_value, *ptr_value, value)
	*ptr_value = 99
	fmt.Printf("After Changes, Ptr Address: %d |  Ptr Value: %d | Value: %d\n", ptr_value, *ptr_value, value)
	change_address := 100000
	ptr_value = &change_address
	fmt.Printf("After Change Address, Ptr Address: %d |  Ptr Value: %d | Value: %d\n", ptr_value, *ptr_value, value)
}

func pointer_nil() {
	var n *int
	fmt.Printf("Nil 0 Address: %d\n", n)

	var n1 *int = nil
	fmt.Printf("Nil 1 Address: %d\n", n1)
}

func pointersExample() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
