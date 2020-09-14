package test

import (
	"fmt"
	tbox "github.com/redrain-wang/golang-toolbox"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println(tbox.GetType("hello world"))
	fmt.Println(tbox.StringToInt("890120"))
	fmt.Println(tbox.IntToString(0b100))    //返回的的是十进制4
	fmt.Println(tbox.StringToBool("true"))  //只有true或者1才会返回true其余都是false
	fmt.Println(tbox.StringToBool("1"))     //只有true或者1才会返回true其余都是false
	fmt.Println(tbox.StringToBool("false")) //除了true和1其余都是false
	fmt.Println(tbox.StringToFloat64("33"))
	fmt.Printf("%v, %T\n", tbox.FloatToString(23.63), tbox.FloatToString(23.63))
	fmt.Println(tbox.StringContain("hello world hello golang", "world"))
	fmt.Println(tbox.StringCounts("hello world hello golang", "o"))    //4
	fmt.Println(tbox.StringIndex("hello world hello golang", "w"))     //6, 不存在返回-1
	fmt.Println(tbox.StringLastIndex("hello world hello golang", "g")) //23, 不存在返回-1
	fmt.Printf("%d,%d\n", len("  golang  "), len(tbox.StringTrim("aagolangaa", "aa")))
	fmt.Printf("%d,%d\n", len("  golang  "), len(tbox.StringTrimLeft("  golang  ")))
	fmt.Printf("%d,%d\n", len("  golang  "), len(tbox.StringTrimRight("  golang  ")))
	fmt.Printf("%d,%d\n", len("fffgolangfff"), len(tbox.StringTrimPrefix("fffgolangfff", "fff"))) //去掉开头的fff
	fmt.Println(tbox.StringCompare("golang", "golang"))                                           // true
	fmt.Println(tbox.StringICompare("golang", "GOlang"))                                          // true
	fmt.Println(tbox.StringField("hello,   golang,  hello, 世 界"))
	fmt.Println(tbox.StringPrefix("hello 世界！！", "hello")) //true
	fmt.Println(tbox.StringSuffix("hello 世界！！", "！"))     //true
	strSlice := []string{"hello", "世界"}
	fmt.Println(tbox.StringJoin(strSlice, ",")) // hello,世界
	sliceRet := tbox.StringSlice("hello,world,世界", ",")
	fmt.Println(sliceRet[2])                                    //世界
	fmt.Printf("%s%s\n", "ba", tbox.StringRepeat("na", 2))      //banana
	fmt.Println(tbox.StringReplaceAll("hello world", "o", "z")) //hellz wzrld (golang 1.12 or latest)
	fmt.Println(tbox.StringReplace("hello world", "o", "z", 1)) // hellz world
	fmt.Println(tbox.StringUcFirst("hello world"))              // Hello World
	fmt.Println(tbox.StringToUpper("hello"))
	fmt.Println(tbox.StringToLower("HELLO"))
	fmt.Println(tbox.StringValidUTF8("hello", "*")) // (golang 1.13 or latest)
	fmt.Println(tbox.StringLen("你好世界"))
	fmt.Println(tbox.StringFormatInt(20000000, 2)) // 1001100010010110100000000
}
