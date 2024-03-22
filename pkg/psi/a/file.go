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

func (mg MyGreeter) Demo(a string) string {
	return "Demo not in Greeter interface: " + a
}

func main() {
	b.MethodHello()

	myStruct := b.MyStruct{Name: "moqi"}
	myStruct.FunctionHello()

	slice1 := []b.MySliceStruct{{Name: "moqi"}}
	fmt.Println(slice1)

	map1 := map[string]b.MyMapStruct{
		"moqi": {Name: "moqi"},
	}
	fmt.Println(map1)

	pointStruct := &b.MyPointerStruct{Name: string("moqi")}
	fmt.Printf("%+v\n", pointStruct)
}
