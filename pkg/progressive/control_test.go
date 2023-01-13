package progressive

import (
	"fmt"
	"testing"
)

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
