package progressive

import (
	"fmt"
	"testing"
)

func TestSwitchCase(t *testing.T) {
	switch x, y := 1, 2; x + y {
	case 3:
		a := 1
		fmt.Println("case1: a = ", a)
		fallthrough
	case 10:
		a := 5
		fmt.Println("case2: a =", a)
		fallthrough
	default:
		a := 7
		fmt.Println("default case: a =", a)
	}
}

func TestSelectCase(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int, 1)
	c2 <- 11

	select {
	case c1 <- 1:
		fmt.Println("SendStmt case has been chosen")
	case i := <-c2:
		_ = i
		fmt.Println("RecvStmt case has been chosen")
	default:
		fmt.Println("default case has been chosen")
	}
}
