// Package gjson
// @author moqi
// On 2022/10/18 18:26:30
// see: https://github.com/tidwall/gjson
package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func TestGjson1(t *testing.T) {
	value := gjson.Get(json, "name.last")
	fmt.Println("value.String() -> ", value.String())
}
