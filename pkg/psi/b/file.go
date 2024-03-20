package b

import "fmt"

type MyStruct struct {
	Name string
}

func (m MyStruct) FunctionHello() {
	fmt.Println("FunctionHello from b function.")
}

func MethodHello() {
	fmt.Println("MethodHello from b method.")
}

func (m MyStruct) FunctionHelloXXXYYY() {
	fmt.Println("FunctionHello from b function xxxyyy.")
}

func (m *MyStruct) FunctionHelloXXXYYYMMM() {
	fmt.Println("FunctionHello from b function xxxyyy.")
}
