package toolbox

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// float64保留两位小数 四舍五入 返回字符串
func Float64TwoPoint(number float64) string {
	if number == 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", number)
}

// 字符串数字转float64
func StringNumberTransToFloat(numberString string) float64 {
	if numberString == "" {
		return 0
	}
	number, _ := strconv.ParseFloat(numberString, 64)
	return number
}

// float64保留两位小数四舍五入，返回float
func Float64TwoPointFloat(number float64) float64 {
	if number == 0 {
		return 0
	}
	numberString := fmt.Sprintf("%.2f", number)
	number, _ = strconv.ParseFloat(numberString, 64)
	return number
}

// float64保留指定位数，返回float
func Float64Round(number float64, n int) float64 {
	numberString := fmt.Sprintf("%0."+strconv.Itoa(n)+"f", number)
	number, _ = strconv.ParseFloat(numberString, 64)
	return number
}

// 十进制int 转换为二进制int
func DecimalToBinary(number int64) int64 {
	numberString := fmt.Sprintf("%d", number)
	number, _ = strconv.ParseInt(numberString, 10, 64)
	return number
}

// 将分转化为元
func GetYuanByFen(num int) float64 {
	return Float64TwoPointFloat(float64(num) / 100)
}

// 将元转化为分
func GetFenByYuan(num float64) int {
	return int(num * 100)
}

// 获取指定范围的随机数字
func GetRandRangeNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
