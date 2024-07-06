package main

import "fmt"

func main() {
	fmt.Println("------- Make -------")

	a := make([]int, 0)
	aptr := &a

	fmt.Println("Pointer == nil: ", *aptr == nil)
	fmt.Printf("Pointer value: %p\n\n", *aptr)

	// composite literal

	fmt.Println("------- Composite Literal -------")
	b := []int{}
	bPtr := &b
	fmt.Println("pointer == nil: ", *bPtr == nil)
	fmt.Printf("pointer value: %p\n\n", *bPtr)

}
