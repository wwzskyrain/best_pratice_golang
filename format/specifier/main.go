package main

import "fmt"

type Person struct {
	age  int
	name string
}

func (p *Person) String() string {
	// 第一次遇到go中的【stack overflow】哈哈
	// fatal error: stack overflow
	// return fmt.Sprintf("%+v", p)
	return "person-string"
}

// PersonSlice 难道要给所有自定义struct都实现其对应的slice么？泛型来解决？
type PersonSlice []*Person

func (pSlice PersonSlice) String() string {
	out := ""
	for _, pp := range pSlice {
		out = out + fmt.Sprintf("%+v ,", *pp)
	}
	return out
}

func main() {

	fmt.Printf("success=%t \n", true)
	p := Person{
		32, "wang-yueyi",
	}
	fmt.Printf("%+v \n", p)
	fmt.Printf("%T \n", p)

	persons := make([]*Person, 1)
	persons = append(persons, &Person{1, "name1"})
	persons = append(persons, &Person{2, "name2"})
	persons = append(persons, &Person{3, "name3"})
	fmt.Println(persons)
	fmt.Println(PersonSlice(persons))
}
