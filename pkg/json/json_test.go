// 如何在Golang中正确序列化JSON字符串
// see: https://blog.csdn.net/Haikuotiankong11111/article/details/109646269
package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func TestJson1(t *testing.T) {
	in := `{"firstName":"John","lastName":"Dow"}`

	bytes, err := json.Marshal(in)
	if err != nil {
		_ = fmt.Errorf("marshal error: %s", err.Error())
	}

	var p Person
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		_ = fmt.Errorf("unmarshal error: %s", err.Error())
	}

	fmt.Printf("%+v", p)
}

func TestJson2(t *testing.T) {
	bytes, err := json.Marshal(Person{
		FirstName: "John",
		LastName:  "Dow",
	})
	if err != nil {
		_ = fmt.Errorf("marshal error: %s", err.Error())
	}

	fmt.Println(string(bytes))
}

func TestJson3(t *testing.T) {
	bytes, err := json.Marshal(Person{
		FirstName: "John",
		LastName:  "Dow",
	})
	if err != nil {
		_ = fmt.Errorf("marshal error: %s", err.Error())
	}

	fmt.Println(string(bytes))

	in := `{"firstName":"John","lastName":"Dow"}`

	bytes, err = json.Marshal(in)
	if err != nil {
		_ = fmt.Errorf("marshal error: %s", err.Error())
	}

	fmt.Println(string(bytes))
}

func TestJson4(t *testing.T) {
	in := `{"firstName":"John","lastName":"Dow"}`

	rawIn := json.RawMessage(in)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		_ = fmt.Errorf("raw message error: %s", err.Error())
	}

	var p Person
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		_ = fmt.Errorf("unmarshal error: %s", err.Error())
	}

	fmt.Printf("%+v", p)
}

func TestJson5(t *testing.T) {
	in := `{"firstName":"John","lastName":"Dow"}`
	bytes := []byte(in)

	var p Person
	err := json.Unmarshal(bytes, &p)
	if err != nil {
		_ = fmt.Errorf("unmarshal error: %s", err.Error())
	}

	fmt.Printf("%+v", p)
}
