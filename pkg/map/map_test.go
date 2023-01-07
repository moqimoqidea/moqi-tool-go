package _map

import (
	"fmt"
	"golang.org/x/exp/maps"
	"testing"
)

func TestMapCopy(t *testing.T) {
	src := map[int]string{200: "foo", 300: "bar"}
	dest := map[int]string{}

	maps.Copy(dest, src)
	fmt.Println("dest ->", dest) // map[200:foo 300:bar]

	dest2 := map[int]string{200: "will be overwritten"}
	maps.Copy(dest2, src)
	fmt.Println("dest2 ->", dest2) // map[200:foo 300:bar]
	fmt.Println(dest)              // map[200:foo 300:bar]
}
