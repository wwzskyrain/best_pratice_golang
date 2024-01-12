package main

import "fmt"

func main() {

	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = i * 10
	}
	fmt.Println(s)

	s = s[:5]
	fmt.Println(s)
}
