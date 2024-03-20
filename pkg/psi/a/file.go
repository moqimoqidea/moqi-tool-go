package a

import (
	"fmt"
	"github.com/moqimoqidea/moqi-tool-go/pkg/psi/b"
)

// MyGreeter 结构体类型
type MyGreeter struct{}

// Hello 方法实现了 b 包中 Greeter 接口的 Hello 方法
func (mg MyGreeter) Hello() string {
	return "Hello from package a"
}

func (mg MyGreeter) HelloWithParameter(a string) string {
	return "Hello from package a with parameter: " + a
}

func main() {
	b.MethodHello()

	myStruct := b.MyStruct{Name: "moqi"}
	myStruct.FunctionHello()

	myStruct.FunctionHelloXXXYYYMMM()

	mg := MyGreeter{}
	greeting := b.Greet(mg) // 使用 b 包的 Greet 函数
	fmt.Println(greeting)
	greetingWithParameter := b.GreetWithParameter(mg, "moqi") // 使用 b 包的 GreetWithParameter 函数
	fmt.Println(greetingWithParameter)
}
