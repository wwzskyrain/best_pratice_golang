package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	// demoFromGPT4()

	debugJsonGet()

}

func debugJsonGet() {

	js := generateJsonObject()
	// 每个v底层都有一个type属性的
	if v, err := js.Get("name").String(); err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(v)
	}
}

func generateJsonObject() simplejson.Json {
	// 我们要先有一个 JSON 字符串
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "isMarried": true,
        "interests": ["Reading", "Cycling", "Hiking"]
    }`
	// 然后使用 simplejson.NewJson 方法解析 JSON 字符串
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err)
	}
	return *js
}

// 这是一个从gpt4直接copy出来的代码，真的很好用
func demoFromGPT4() {

	// 假设我们有一个 JSON 字符串
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "isMarried": true,
        "interests": ["Reading", "Cycling", "Hiking"]
    }`

	// 使用 simplejson.NewJson 方法解析 JSON 字符串
	js, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		panic(err)
	}

	// 获取并打印 name 字段
	name, err := js.Get("name").String()
	if err != nil {
		panic(err)
	}
	fmt.Println("Name:", name)

	// 获取并打印 age 字段
	age, err := js.Get("age").Int()
	if err != nil {
		panic(err)
	}
	fmt.Println("Age:", age)

	// 获取并迭代打印 interests 数组
	interests, err := js.Get("interests").Array()
	if err != nil {
		panic(err)
	}
	fmt.Println("Interests:")
	for i, val := range interests {
		fmt.Printf(" - Interest #%d: %s\n", i+1, val)
	}

	// 修改 age 字段
	js.Set("age", 31)

	// 添加一个新字段
	js.Set("city", "New York")

	// 将修改后的对象转换回 JSON 字符串
	newJsonStr, err := js.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Modified JSON:", string(newJsonStr))
}
