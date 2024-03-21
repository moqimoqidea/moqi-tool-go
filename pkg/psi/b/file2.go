package b

import "fmt"

func (m MyStruct) FunctionHelloXXX() {
	fmt.Println("FunctionHello from b function xxx.")
}

func (m MyStruct) functionInside() {
	fmt.Println("functionInside from b MyStruct.")
}
