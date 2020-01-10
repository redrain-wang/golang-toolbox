package toolbox

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 获取变量类型
func GetType(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

// 字符串转换为int
func StringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}

// int转换为string
func IntToString(number int) string {
	str := strconv.Itoa(number)
	return str
}

// int转换为 string 类型 并指定进制
func StringFormatInt(number int64, base int) string {
	if base < 2 || base > 36 {
		panic("base range 2 and 36")
	}
	return strconv.FormatInt(number, base)
}

// string转换为bool
func StringToBool(str string) bool {
	result, _ := strconv.ParseBool(str)
	return result
}

// string 转float64
func StringToFloat64(str string) float64 {
	result, _ := strconv.ParseFloat(str, 64)
	return result
}

// float转换为 string 保留两位小数
func FloatToString(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

// float转换为字符串 保留指定位数 会自动四舍五入
func FormatFloat(number float64, prec int) string {
	return strconv.FormatFloat(number, 'f', prec, 64)
}

// 查找某个字符或串是否包含在给定的字符串里
func StringContain(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}

// 查找一个字符串在某个字符串中出现的次数
func StringCounts(haystack, needle string) int {
	return strings.Count(haystack, needle)
}

// 查找字符串第一出现的位置 不存在返回-1
func StringIndex(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// 查找字符串最后一次出现的位置 不存在返回 -1
func StringLastIndex(haystack, needle string) int {
	return strings.LastIndex(haystack, needle)
}

// 去除字符两边空格
func StringTrim(str string, tag string) string {
	return strings.Trim(str, tag)
}

// 替换字符串两边空格
func StringTrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// 去除左侧空格
func StringTrimLeft(str string) string {
	return strings.TrimLeft(str, " ")
}

// 去除右侧空格
func StringTrimRight(str string) string {
	return strings.TrimRight(str, " ")
}

// 从开头去除指定的字符
func StringTrimPrefix(str, prefix string) string {
	return strings.TrimPrefix(str, prefix)
}

// 判断两个字符串是否完全相等 区分大小写
func StringCompare(str1, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

// 判断两个字符串是否完全相等 不区分大小写
func StringICompare(str1, str2 string) bool {
	str1 = strings.ToUpper(str1)
	str2 = strings.ToUpper(str2)
	return StringCompare(str1, str2)
}

// 将给定的字符串用空白符分割为切片
func StringField(str string) []string {
	return strings.Fields(str)
}

// 判断给定字符串是否是以指定的字符开头的
func StringPrefix(haystack, prefix string) bool {
	return strings.HasPrefix(haystack, prefix)
}

// 判断给定字符串是否以指定字符串结尾
func StringSuffix(haystack, suffix string) bool {
	return strings.HasSuffix(haystack, suffix)
}

// 将slice转换为string
func StringJoin(v []string, sep string) string {
	return strings.Join(v, sep)
}

//string 转化为slice
func StringSlice(str, sep string) []string {
	return strings.Split(str, sep)
}

// 将指定字符串重复出现n次
func StringRepeat(str string, number int) string {
	return strings.Repeat(str, number)
}

// 将指定字串替换为给定的串 从头替换到尾
func StringReplaceAll(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// 将指定字串替换为给定的串 替换指定个数
func StringReplace(str, old, new string, number int) string {
	return strings.Replace(str, old, new, number)
}

// 首字母大写
func StringUcFirst(str string) string {
	return strings.Title(str)
}

// 转换为大写
func StringToUpper(str string) string {
	return strings.ToUpper(str)
}

// 转换为小写
func StringToLower(str string) string {
	return strings.ToLower(str)
}

// 将无效的UTF8字节序列替换为 给定的字符串
func StringValidUTF8(str1, str2 string) string {
	return strings.ToValidUTF8(str1, str2)
}

// 获取utf8字符串长度
func StringLen(str string) int {
	return utf8.RuneCountInString(str)
}
