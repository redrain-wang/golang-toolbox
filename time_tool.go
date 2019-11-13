package time_tool

import (
	"fmt"
	"testing"
	"time"
)

// 获取当前时间 time.Time类型
func GetNow() time.Time {
	return time.Now()
}

// 获取当前日期和时间 格式 2019-01-01 10：10：10
func GetNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前日期 格式 2019-01-01
func GetTodayDate() string {
	return time.Now().Format("2006-01-02")
}

// 获取当前时间
func GetCurrentTime() string {
	return time.Now().Format("15:04:05")
}

// 获取当前时间戳 秒级
func GetTimestampSecond() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳 毫秒级
func GetTimestampMill() int64 {
	return time.Now().UnixNano() / 1e6
}

// 获取当前时间戳 微秒级
func GetTimestampMico() int64 {
	return time.Now().UnixNano() / 1e3
}

// 获取当前时间戳 纳秒级
func GetTimeStampNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前是一年中的第几天
func GetThDayInYear() int {
	return time.Now().YearDay()
}

// 获取当前是一月中的第几天
func GetThDayInMonth() int {
	return time.Now().Day()
}

// 获取当前星期 英文全拼
func GetWeekEnglish() string {
	return time.Now().Weekday().String()
}

// 获取当前星期 阿拉伯数字
func GetWeekInt() int {
	return int(time.Now().Weekday())
}

// 获取当前星期 中文描述
func GetWeekChina() string {
	weekNum := int(time.Now().Weekday())
	switch weekNum {
	case 0:
		return "星期日"
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	case 6:
		return "星期六"
	}

	return ""
}

// 获取当前本年度第一天的星期
func GetFirstDayWeekOfYear() int {
	return int(GetNow().AddDate(0, 0, -GetNow().YearDay()+1).Weekday())
}

// 获取本年度第一周有几天
func GetFirstWeekDayCountsOfYear() int {
	if GetFirstDayWeekOfYear() != 0 {
		return 7 - GetFirstDayWeekOfYear() + 1
	}
	return 1
}

// 获取当前是本年度第几周
func GetWeekThOfYear() int {
	yearTh := GetNow().YearDay()
	if yearTh < GetFirstWeekDayCountsOfYear() {
		return 1
	}
	return (yearTh-GetFirstWeekDayCountsOfYear())/7 + 2
}

// 获取本月第一天的日期
func GetFirstDayDateOfMonth() string {
	return time.Now().AddDate(0, 0, -time.Now().Day()+1).Format("2006-01-02")
}

// 获取本月第一天的星期
func GetFirstDayWeekOfMonth() int {
	return int(time.Now().AddDate(0, 0, -time.Now().Day()+1).Weekday())
}

// 获取本月第一周有几天
func GetFirstWeekDayCountsOfMonth() int {
	if GetFirstDayWeekOfMonth() != 0 {
		return 7 - GetFirstDayWeekOfMonth() + 1
	}
	return 1
}

// 获取当前是本月第几周
func GetWeekThOfMonth() int {
	if GetNow().Day() < GetFirstWeekDayCountsOfMonth() {
		return 1
	}
	return (GetNow().Day()-GetFirstWeekDayCountsOfMonth())/7 + 2
}

func TestTime(t *testing.T) {
	fmt.Println(GetNowString())
	fmt.Println(GetTodayDate())
	fmt.Println(GetCurrentTime())
	fmt.Println(GetTimestampSecond())
	fmt.Println(GetTimestampMill())
	fmt.Println(GetTimestampMico())
	fmt.Println(GetTimeStampNano())
	fmt.Println(GetThDayInYear())
	fmt.Println(GetThDayInMonth())
	fmt.Println(GetWeekEnglish())
	fmt.Println(GetWeekInt())
	fmt.Println(GetWeekChina())
	fmt.Println(GetFirstDayWeekOfYear())
	fmt.Println(GetFirstWeekDayCountsOfYear())
	fmt.Println(GetWeekThOfYear())
	fmt.Println(GetFirstDayDateOfMonth())
	fmt.Println(GetFirstDayWeekOfMonth())
	fmt.Println(GetFirstWeekDayCountsOfMonth())
	fmt.Println(GetWeekThOfMonth())
}
