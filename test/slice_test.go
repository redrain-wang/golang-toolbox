package test

import (
	tbox "github.com/redrain-wang/golang-toolbox"
	"testing"
)

func TestSlice(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(tbox.InArray(slice1, 2))  // true
	t.Log(tbox.InArray(slice1, 11)) // false
}

