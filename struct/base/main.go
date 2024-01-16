package main

import "fmt"

type Person struct {
	age  int
	name string
}

func main() {
	var p Person
	if &p != nil { // 这里必须要用&p，不能用p的，否则编译报错：invalid operation: p == nil (mismatched types Person and untyped nil)
		fmt.Println("p is not nil")
	}
	fmt.Println("p：", p)

}
