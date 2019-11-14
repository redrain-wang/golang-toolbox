package test

import (
	"fmt"
	tt "github.com/redrain-wang/golang-toolbox/timeTool"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(tt.GetNowString())
	fmt.Println(tt.GetTodayDate())
	fmt.Println(tt.GetCurrentTime())
	fmt.Println(tt.GetTimestampSecond())
	fmt.Println(tt.GetTimestampMill())
	fmt.Println(tt.GetTimestampMico())
	fmt.Println(tt.GetTimeStampNano())
	fmt.Println(tt.GetThDayInYear())
	fmt.Println(tt.GetThDayInMonth())
	fmt.Println(tt.GetWeekEnglish())
	fmt.Println(tt.GetWeekInt())
	fmt.Println(tt.GetWeekChina())
	fmt.Println(tt.GetFirstDayWeekOfYear())
	fmt.Println(tt.GetFirstWeekDayCountsOfYear())
	fmt.Println(tt.GetWeekThOfYear())
	fmt.Println(tt.GetFirstDayDateOfMonth())
	fmt.Println(tt.GetFirstDayWeekOfMonth())
	fmt.Println(tt.GetFirstWeekDayCountsOfMonth())
	fmt.Println(tt.GetWeekThOfMonth())
}
