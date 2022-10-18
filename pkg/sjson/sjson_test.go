// Package sjson
// @author moqi
// On 2022/10/18 18:55:30
// see: https://github.com/tidwall/sjson
package gjson

import (
	"fmt"
	"github.com/tidwall/sjson"
	"testing"
)

const (
	simpleJson = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
)

func TestGjson1(t *testing.T) {
	value, _ := sjson.Set(simpleJson, "name.last", "Anderson")
	fmt.Println("value -> ", value)
}
