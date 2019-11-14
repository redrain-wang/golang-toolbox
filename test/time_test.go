package test

import (
	"fmt"
	tbox "github.com/redrain-wang/golang-toolbox"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(tbox.GetNowString())
	fmt.Println(tbox.GetTodayDate())
	fmt.Println(tbox.GetCurrentTime())
	fmt.Println(tbox.GetTimestampSecond())
	fmt.Println(tbox.GetTimestampMill())
	fmt.Println(tbox.GetTimestampMico())
	fmt.Println(tbox.GetTimeStampNano())
	fmt.Println(tbox.GetThDayInYear())
	fmt.Println(tbox.GetThDayInMonth())
	fmt.Println(tbox.GetWeekEnglish())
	fmt.Println(tbox.GetWeekInt())
	fmt.Println(tbox.GetWeekChina())
	fmt.Println(tbox.GetFirstDayWeekOfYear())
	fmt.Println(tbox.GetFirstWeekDayCountsOfYear())
	fmt.Println(tbox.GetWeekThOfYear())
	fmt.Println(tbox.GetFirstDayDateOfMonth())
	fmt.Println(tbox.GetLastDayDateOfMonth())
	fmt.Println(tbox.GetDayCountsOfMonth())
	fmt.Println(tbox.GetFirstDayWeekOfMonth())
	fmt.Println(tbox.GetFirstWeekDayCountsOfMonth())
	fmt.Println(tbox.GetWeekThOfMonth())
}

func TestSlice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(tbox.InArray(slice1, 2))  // true
	t.Log(tbox.InArray(slice1, 11)) // false
}
