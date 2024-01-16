package main

import "fmt"

type Person struct {
	age  int
	name string
}

func main() {

	fmt.Printf("success=%t \n", true)
	p := Person{
		32, "wang-yueyi",
	}
	fmt.Printf("%+v \n", p)
	fmt.Printf("%T \n", p)
}
