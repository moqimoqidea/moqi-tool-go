// Package handle
// @author moqi
// On 2021/9/5 23:06:59
package handle

import (
	"testing"
)

func TestTriggerError(t *testing.T) {
	want := "base error context: base error"
	if got := TriggerError().Error(); got != want {
		t.Errorf("TestTriggerError() = %q, want %q", got, want)
	}
}
