package b

import "fmt"

// Greeter 接口定义了一个方法签名 Hello
type Greeter interface {
	Hello() string
	HelloWithParameter(a string) string
}

// Greet 函数接受实现了 Greeter 接口的任何类型的实例
func Greet(g Greeter) string {
	return g.Hello()
}

func GreetWithParameter(g Greeter, a string) string {
	return g.HelloWithParameter(a)
}

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
