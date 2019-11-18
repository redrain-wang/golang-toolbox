package toolbox

import (
	"fmt"
	"strconv"
)

// float64保留两位小数 四舍五入 返回字符串
func Float64TwoPoint(number float64) string{
	if number == 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", number)
}

// 字符串数字转float64
func StringNumberTransToFloat(numberString string) float64{
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
	number , _ = strconv.ParseFloat(numberString, 64)
	return number
}

//