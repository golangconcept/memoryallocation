package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

func main() {
	n := new(Name)
	n.FirstName = "hi"

	fmt.Println(n.FirstName)
}
