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

func TestTime0(t *testing.T) {
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	fmt.Println(currentDate)
}
func TestTime(t *testing.T) {
	currentTime := time.Now()

	fmt.Println("Show Current Time in String: ", currentTime.String())
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))
	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println("ShortNum Width : ", currentTime.Format("2006-01-02"))
	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))
	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))
	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))
	fmt.Println("LongWidth : ", currentTime.Format("2006-January-02"))
	fmt.Println("ShortWidth : ", currentTime.Format("2006-Jan-02"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))
}

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
