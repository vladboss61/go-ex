package main

import (
	"fmt"
)

type DataAllocation struct {
	v *int;
}

// new operator allocates memory in the heap and returns pointer to that memory.
func new_example() {
	i := new(int32)
	fmt.Printf("1 Address: %d | Value: %d\n", i, *i)
	*i = 55

	fmt.Printf("2 Address: %d | Value: %d\n", i, *i)

	var s *DataAllocation = &DataAllocation{
		v: new(int),
	}

	if s.v == nil {
		fmt.Printf("3 Address: %d | Value: nil", s.v)
		
	} else {
		fmt.Printf("3 Address: %d | Value: %v\n", s.v, *s.v)
	}

	*(s.v) = 25

	fmt.Printf("3 Address: %d | Value: %d\n", s.v, *s.v)
}