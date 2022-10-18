// Package now
// @author moqi
// On 2022/10/18 18:22:11
// see: https://github.com/jinzhu/now
package now

import (
	"fmt"
	"github.com/jinzhu/now"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	fmt.Println("time.Now() ->", time.Now())

	fmt.Println("now.BeginningOfMinute() ->", now.BeginningOfMinute())
	fmt.Println("now.BeginningOfHour() ->", now.BeginningOfHour())
	fmt.Println("now.BeginningOfDay() ->", now.BeginningOfDay())
	fmt.Println("now.BeginningOfWeek() ->", now.BeginningOfWeek())
	fmt.Println("now.BeginningOfMonth() ->", now.BeginningOfMonth())
	fmt.Println("now.BeginningOfQuarter() ->", now.BeginningOfQuarter())
	fmt.Println("now.BeginningOfYear() ->", now.BeginningOfYear())

	fmt.Println("now.EndOfMinute() ->", now.EndOfMinute())
	fmt.Println("now.EndOfHour() ->", now.EndOfHour())
	fmt.Println("now.EndOfDay() ->", now.EndOfDay())
	fmt.Println("now.EndOfWeek() ->", now.EndOfWeek())
	fmt.Println("now.EndOfMonth() ->", now.EndOfMonth())
	fmt.Println("now.EndOfQuarter() ->", now.EndOfQuarter())
	fmt.Println("now.EndOfYear() ->", now.EndOfYear())

	now.WeekStartDay = time.Monday
	fmt.Println("now.EndOfWeek() ->", now.EndOfWeek())
}
