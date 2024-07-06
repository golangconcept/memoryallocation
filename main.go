package main

import "fmt"

func main() {
	fmt.Printf("Hello, Jai Pal\n")

	a := make([]int, 0)
	aptr := &a

	fmt.Println("Pointer == nil: ", *aptr == nil)
	fmt.Println("Pointer value:%p\n\n", *aptr)

}
