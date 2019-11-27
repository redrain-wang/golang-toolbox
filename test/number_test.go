package test

import (
	"fmt"
	tbox "github.com/redrain-wang/golang-toolbox"
	"testing"
)

func TestNumber(t *testing.T) {
	fmt.Println(tbox.Float64TwoPoint(462360.01666666666))
	fmt.Println(tbox.Float64TwoPointFloat(462360.01666666666))
	fmt.Println(tbox.StringNumberTransToFloat("196578"))
	fmt.Println(tbox.DecimalToBinary(10000000))
	fmt.Println(tbox.GetYuanByFen(23453))
	fmt.Println(tbox.GetFenByYuan(22.3554))

}