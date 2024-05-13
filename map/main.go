package main

import "fmt"

func testAppendValueToSliceWhenSliceIsMapValue() {

	m := map[string]int{
		"wang":  18,
		"zhang": 18,
		"li":    16,
	}

	fmt.Printf("%v \n", m)

	age2NameSliceMap := make(map[int][]string)

	for name, age := range m {
		names, ok := age2NameSliceMap[age]
		if !ok {
			names = make([]string, 0)
		}
		names = append(names, name)
		age2NameSliceMap[age] = names
	}

	fmt.Printf("age2NameSliceMap = %v \n", age2NameSliceMap)

}

func main() {
	testAppendValueToSliceWhenSliceIsMapValue()
	Test_Npe()
}

func Test_Npe() {
	names := []string{"1", "2"}
	var siteIdMap = make(map[string]bool)
	for _, name := range names {
		siteIdMap[name] = true
	}
}
