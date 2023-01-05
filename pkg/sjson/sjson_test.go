// Package sjson
// @author moqi
// On 2022/10/18 18:55:30
// see: https://github.com/tidwall/sjson
package sjson

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/sjson"
	"testing"
)

const (
	simpleJson = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
)

func TestSjson1(t *testing.T) {
	value, _ := sjson.Set(simpleJson, "name.last", "Anderson")
	fmt.Println("value -> ", value)
}

func TestSjson2(t *testing.T) {
	tempMap := make(map[string]string)
	tempMap["k1"] = "v1"
	tempMap["k2"] = "v2"
	tempMap["k3"] = "v3"

	jsonStr, _ := json.Marshal(tempMap)
	fmt.Println("jsonStr -> ", jsonStr)
	value, _ := sjson.SetRaw(simpleJson, "mapData", string(jsonStr))
	fmt.Println("value -> ", value)
}
