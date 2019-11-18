package toolbox

import (
	"fmt"
	"time"
)

// 获取当前时间 time.Time类型
func GetNow() time.Time {
	return time.Now()
}

// 将给定的完整日期时间字符串转化为golang的Time
func TransTimeStringToGoTime(timeString string) time.Time {
	goTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err == nil {
		return goTime
	}
	return GetNow()
}

// 以给定的格式字符串返回go的Time类型的时间
func TransTimeStringByFormatToGoTime(formatString, timeString string) time.Time {
	goTime, err := time.Parse(formatString, timeString)
	if err == nil {
		return goTime
	}
	return GetNow()
}

// 获取当前年
func GetYear() int {
	return GetNow().Year()
}

// 获取定日期得年份
func GetYearByFormatTime(formatString, timeString string) int {
	return TransTimeStringByFormatToGoTime(formatString, timeString).Year()
}

//获取当前年 2019年
func GetYearString() string {
	return fmt.Sprintf("%d年", GetNow().Year())
}

// 获取定日期得年份
func GetYearStringByFormatTime(formatString, timeString string) string {
	return fmt.Sprintf("%d年", GetYearByFormatTime(formatString, timeString))
}

// 获取当前月 英文全拼
func GetMonthString() string {
	return GetNow().Month().String()
}

// 获取当前月 阿拉伯数字
func GetMonthNumber() int {
	return int(GetNow().Month())
}

// 获取当前月的中文描述
func GetMonthChina() string {
	return GetMonthChinaByNum(GetMonthNumber())
}

// 获取给定月的中文描述 一月 二月
func GetMonthChinaByNum(num int) string {
	if num > 12 || num <= 0 {
		return "month error"
	}
	monthMap := make(map[int]string, 12)
	monthMap[1] = "一月"
	monthMap[2] = "二月"
	monthMap[3] = "三月"
	monthMap[4] = "四月"
	monthMap[5] = "五月"
	monthMap[6] = "六月"
	monthMap[7] = "七月"
	monthMap[8] = "八月"
	monthMap[9] = "九月"
	monthMap[10] = "十月"
	monthMap[11] = "十一月"
	monthMap[12] = "十二月"
	return monthMap[num]
}

// 获取当前日期的日
func GetDay() int {
	return GetNow().Day()
}

// 获取当前时间的小时
func GetHour() int {
	return GetNow().Hour()
}

// 获取当前时间的分钟
func GetMinute() int {
	return GetNow().Minute()
}

// 获取当前时间秒
func GetSecond() int {
	return GetNow().Second()
}

// 获取当前纳秒
func GetNanoSecond() int {
	return GetNow().Nanosecond()
}

// 获取当前毫秒
func GetMillSecond() int {
	return GetNanoSecond() / 1e6
}

// 获取当前微秒
func GetMicroSecond() int {
	return GetNanoSecond() / 1e3
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
	return time.Date(GetNow().Year(), GetNow().Month(), 1, 0, 0, 0, 0, GetNow().Location()).Format("2006-01-02")
}

// 获取本月最后一天的日期
func GetLastDayDateOfMonth() string {
	return time.Date(GetNow().Year(), GetNow().Month(), 1, 0, 0, 0, 0, GetNow().Location()).AddDate(0, 1, -1).Format("2006-01-02")
}

// 获取本月有多少天
func GetDayCountsOfMonth() int {
	return time.Date(GetNow().Year(), GetNow().Month(), 1, 0, 0, 0, 0, GetNow().Location()).AddDate(0, 1, -1).Day()
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

// 将当前的时间加指定的天数
func GetCurrentTimeAddDays(days int) string {
	return GetNow().AddDate(0,0,1).Format("2006-01-02 15:04:05")
}

func GetCurrentTimeAddMonths(months int) string {
	return GetNow().AddDate(0, months, 0).Format("2006-01-02 15:04:05")
}

// 将当前时间加指定时间返回秒数
func GetCurrentTimeAddSeconds(seconds int64) int64 {
	return GetNow().Unix() + seconds
}

// 将当前时间添加指定秒数返回时间字符串
func GetCurrentTimeAddSecondsString(seconds int64) string {
	return GetTimeStringByUnixTime(GetCurrentTimeAddSeconds(seconds))
}

// 将给定时间戳转化为时间
func GetTimeStringByUnixTime(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

// 获取两个时间相差的秒数 time1 - time2
func GetTimeDiffSecond(time1, time2 string) int64 {
	return TransTimeStringToGoTime(time1).Unix() - TransTimeStringToGoTime(time2).Unix()
}

// 获取两个时间相差分钟数
func GetTimeDiffMinute(time1, time2 string) float64 {
	return Float64TwoPointFloat(float64(GetTimeDiffSecond(time1, time2)) / 60)
}

// 是否是星期日
func IsSunday() bool{
	if int(GetNow().Weekday()) == 0 {
		return true
	}
	return false
}

// 是否是星期一
func IsMonday() bool{
	if int(GetNow().Weekday()) == 1 {
		return true
	}
	return false
}

// 是否星期二
func IsTuesday() bool{
	if int(GetNow().Weekday()) == 2 {
		return true
	}
	return false
}
// 是否星期三
func IsWednesday() bool {
	if int(GetNow().Weekday()) == 3 {
		return true
	}
	return false
}

// 是否星期四
func IsThursday() bool {
	if int(GetNow().Weekday()) == 4 {
		return true
	}
	return false
}

// 是否星期五
func IsFriday() bool {
	if int(GetNow().Weekday()) == 5 {
		return true
	}
	return false
}

// 是否星期六
func IsSaturday() bool{
	if int(GetNow().Weekday()) == 6 {
		return true
	}
	return false
}

// 是否是当前年
func IsCurrentYear(date string) bool {
	if GetNow().Year() == TransTimeStringToGoTime(date).Year() {
		return true
	}
	return false
}

// 是否是去年
func IsLastYear(date string) bool {
	if GetNow().Year() -1  == TransTimeStringToGoTime(date).Year() {
		return true
	}
	return false
}

// 是否是明年
func IsNextYear(date string) bool {
	if GetNow().Year() + 1  == TransTimeStringToGoTime(date).Year() {
		return true
	}
	return false
}