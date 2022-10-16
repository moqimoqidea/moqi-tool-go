package funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

func TestExist(t *testing.T) {
	// 判断任意类型
	fmt.Println("str->", funk.Contains([]string{"a", "b"}, "a"))
	// int 类型
	fmt.Println("int->", funk.ContainsInt([]int{1, 2}, 1))
}

/**输出
=== RUN   TestExist
str-> true
int-> true
--- PASS: TestExist (0.00s)
PASS
*/
