package test

import (
	"testing"

	tbox "github.com/redrain-wang/golang-toolbox"
)

func TestSlice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(tbox.InArray(slice1, 2))  // true
	t.Log(tbox.InArray(slice1, 11)) // false
	t.Log(tbox.MaxInt([]int{1,2,3,4,50,6,7}))
	t.Log(tbox.MaxInt64([]int64{100,2,3,4,50,6,7}))
	t.Log(tbox.MaxFloat32([]float32{1,2,3,4,50,60,7}))
	t.Log(tbox.MaxFloat64([]float64{1,2,3,4,50,6,70}))

	t.Log(tbox.MinInt([]int{1,2,3,4,50,6,7}))
	t.Log(tbox.MinInt64([]int64{100,2,3,4,50,6,7}))
	t.Log(tbox.MinFloat32([]float32{1,2,3,4,50,60,7}))
	t.Log(tbox.MinFloat64([]float64{1,2,3,4,50,6,70}))
	arr := []int{1,2,3,4,50,6,70}
	tbox.SortAscInt(arr)
	t.Log(arr)
	tbox.SortDescInt(arr)
	t.Log(arr)


}

