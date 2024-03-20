package a

import "github.com/moqimoqidea/moqi-tool-go/pkg/psi/b"

func main() {
	b.MethodHello()

	myStruct := b.MyStruct{Name: "moqi"}
	myStruct.FunctionHello()

	myStruct.FunctionHelloXXXYYYMMM()
}
