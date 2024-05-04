package main

import "fmt"

func map_static() {
	// Stack allocation
	dictionary := map[string]int{
		"Vlad": 26,
		"Jon":  30,
		"Pen":  15,
	}
	fmt.Printf("Map Static: %v", dictionary)
}

func map_dynamic() {

	// Heap allocation.
	dictionary := make(map[string]int)
	dictionary["Vlad D"] = 26
	dictionary["Tom"] = 28

	fmt.Printf("Map Dynamic: %v\n", dictionary)

	fmt.Println("=============================")

	fmt.Printf("Length: %v\n", len(dictionary))

	fmt.Println("=============================")

	for key, val := range dictionary {
		fmt.Printf("Key: %s | Value: %d\n", key, val)
	}
}
