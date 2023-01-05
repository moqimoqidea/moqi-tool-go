// Package gjson
// @author moqi
// On 2022/10/18 18:26:30
// see: https://github.com/tidwall/gjson
package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

const (
	simpleJson = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	json       = `{
				  "name": {"first": "Tom", "last": "Anderson"},
				  "age":37,
				  "children": ["Sara","Alex","Jack"],
				  "fav.movie": "Deer Hunter",
				  "friends": [
					{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
					{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
					{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
				  ]
				}`
	jsonLine = `{"name": "Gilbert", "age": 61}
				{"name": "Alexa", "age": 34}
				{"name": "May", "age": 57}
				{"name": "Deloise", "age": 44}`
)

func TestGjson8(t *testing.T) {
	fmt.Println(JsonStringToMap(json, "friends.#.first", "friends.#.age"))
}

func TestGjson7(t *testing.T) {
	fmt.Println(gjson.GetMany(json, "friends.#.first", "friends.#.age"))
}

func TestGjson1(t *testing.T) {
	value := gjson.Get(simpleJson, "name.last")
	fmt.Println("value.String() -> ", value.String())
}

func TestGjson2(t *testing.T) {
	fmt.Println(gjson.Get(json, "name.last"))
	fmt.Println(gjson.Get(json, "age"))
	fmt.Println(gjson.Get(json, "children"))
	fmt.Println(gjson.Get(json, "children.#"))
	fmt.Println(gjson.Get(json, "children.1"))
	fmt.Println(gjson.Get(json, "child*.2"))
	fmt.Println(gjson.Get(json, "c?ildren.0"))
	fmt.Println(gjson.Get(json, "fav\\.movie"))
	fmt.Println(gjson.Get(json, "friends.#.first"))
	fmt.Println(gjson.Get(json, "friends.1.last"))
}

func TestGjson3(t *testing.T) {
	fmt.Println(gjson.Get(json, "friends.#(last==\"Murphy\").first"))
	fmt.Println(gjson.Get(json, "friends.#(last==\"Murphy\")#.first"))
	fmt.Println(gjson.Get(json, "friends.#(age>45)#.last"))
	fmt.Println(gjson.Get(json, "friends.#(first%\"D*\").last"))
	fmt.Println(gjson.Get(json, "friends.#(first!%\"D*\").last"))
	fmt.Println(gjson.Get(json, "friends.#(nets.#(==\"fb\"))#.first"))
}

func TestGjson4(t *testing.T) {
	fmt.Println(gjson.Get(json, "children|@reverse"))
	fmt.Println(gjson.Get(json, "children|@reverse|0"))
}

func TestGjson5(t *testing.T) {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		if arg == "lower" {
			return strings.ToLower(json)
		}
		return json
	})

	fmt.Println(gjson.Get(json, "children|@case:upper"))
	fmt.Println(gjson.Get(json, "children|@case:lower|@reverse"))
}

func TestGjson6(t *testing.T) {
	fmt.Println(gjson.Get(jsonLine, "..#"))
	fmt.Println(gjson.Get(jsonLine, "..1"))
	fmt.Println(gjson.Get(jsonLine, "..3"))
	fmt.Println(gjson.Get(jsonLine, "..#.name"))
	fmt.Println(gjson.Get(jsonLine, "..#(name=\"May\")"))
}
