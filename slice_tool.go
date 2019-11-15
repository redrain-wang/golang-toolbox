package toolbox

import "reflect"

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
