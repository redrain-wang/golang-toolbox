package toolbox

import (
	"reflect"
	"sort"
)

// InArray 检测一个值是否在切片中
func InArray(array interface{}, item interface{}) bool {
	values := reflect.ValueOf(array)
	if values.Kind() != reflect.Slice {
		return false
	}

	size := values.Len()
	list := make([]interface{}, size)
	for index := 0; index < size; index++ {
		//转换为Interface类型
		list[index] = values.Index(index).Interface()
	}

	for index := 0; index < len(list); index++ {
		if list[index] == item {
			return true
		}
	}
	return false
}

// 获取切片最大值 int
func MaxInt(array []int) int {
	var maxValue = array[0]
	for _, val := range array {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

// 获取切片最大值 int64
func MaxInt64(array []int64) int64 {
	var maxValue = array[0]
	for _, val := range array {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

// 获取切片最大值 Float32
func MaxFloat32(array []float32) float32 {
	var maxValue = array[0]
	for _, val := range array {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}
// 获取切片最大值 Float64
func MaxFloat64(array []float64) float64 {
	var maxValue = array[0]
	for _, val := range array {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

// 获取切片最小值 int
func MinInt(array []int) int {
	var minValue = array[0]
	for _, val := range array {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}

// 获取切片最小值 int64
func MinInt64(array []int64) int64 {
	var minValue = array[0]
	for _, val := range array {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}

// 获取切片最小值 Float32
func MinFloat32(array []float32) float32 {
	var minValue = array[0]
	for _, val := range array {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}
// 获取切片最小值 Float64
func MinFloat64(array []float64) float64 {
	var minValue = array[0]
	for _, val := range array {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}

// 从小到大排序slice int
func SortAscInt(array []int) {
	sort.Ints(array)
}

// 从大到小排序slice int
func SortDescInt(array []int) {
	sort.Slice(array, func(i,j int) bool{
		return array[i] > array[j]
	})
}

// 从小到大排序 slice int64
func SortAscInt64(array []int64) {
	sort.Slice(array, func(i,j int) bool{
		return array[i] > array[j]
	})
}

// 从大到小排序slice int64
func SortDescInt64(array []int64) {
	sort.Slice(array, func(i,j int) bool{
		return array[i] > array[j]
	})
}

// 从小到大排序 float64
func SortAscFloat64(array []float64) {
	sort.Float64s(array)
}

// 从大到小排序 float64
func SortDescFloat64(array []float64) {
	sort.Slice(array, func(i,j int) bool{
		return array[i] > array[j]
	})
}
