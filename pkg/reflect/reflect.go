package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  25,
	}

	// 获取结构体字段的值
	value := reflect.ValueOf(p)
	nameValue := value.FieldByName("Name")
	ageValue := value.FieldByName("Age")

	fmt.Println("Name:", nameValue.String())
	fmt.Println("Age:", ageValue.Int())
}
