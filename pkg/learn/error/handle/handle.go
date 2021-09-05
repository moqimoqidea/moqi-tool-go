// Package handle
// @author moqi
// On 2021/9/5 22:58:09
package handle

//goland:noinspection GoSnakeCaseUsage
import (
	go_base_errors "errors"
	"github.com/pkg/errors"
)

// TriggerError
// https://github.com/pkg/errors
// 尝试第三方 errors 工具包
func TriggerError() error {
	return errors.Wrap(go_base_errors.New("base error"), "base error context")
}
