// Package moqi_tool_go
// @author moqi
// On 2021/9/5 12:02:08
package moqi_tool_go

import "testing"

func TestTwoSum(t *testing.T) {
	want := 5
	if got := TwoSum(2, 3); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
