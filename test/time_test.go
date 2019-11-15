package test

import (
	"fmt"
	tbox "github.com/redrain-wang/golang-toolbox"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(tbox.GetNowString())
	fmt.Println(tbox.TransTimeStringToGoTime("2019-11-15 11:17:23"))
	fmt.Println(tbox.TransTimeStringByFormatToGoTime("2006-01-02 15:04", "2019-11-15 11:17"))
	fmt.Println(tbox.GetYear())
	fmt.Println(tbox.GetMonthString())
	fmt.Println(tbox.GetMonthNumber())
	fmt.Println(tbox.GetMonthChina())
	fmt.Println(tbox.GetMonthChinaByNum(int(tbox.GetMonthNumber())))
	fmt.Println(tbox.GetDay())
	fmt.Println(tbox.GetHour())
	fmt.Println(tbox.GetMinute())
	fmt.Println(tbox.GetSecond())
	fmt.Println(tbox.GetMillSecond())
	fmt.Println(tbox.GetNanoSecond())
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
