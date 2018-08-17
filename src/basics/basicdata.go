package basics

import (
	"fmt"
	"unsafe"
	"math"
)

//方法名不大写的话 其他包无法识别 var 全局变量 双引号中的任何值都是 Go 中的字符串常量 其实和java一样 const 常量 :=局部变量 byte 是 uint8 的别名。
//rune 是 int32 的别名。
var testData = 12

const testCon = "sb"

func ShowData() {
	var (
		name         = "crush"
		age     int8 = 23
		isMan        = true
		count        = math.Cos(0.333)
		welCome      = getStr(name)
		old          = float32(age)
	)
	testData = 13
	fmt.Println(name, age, isMan, count, welCome, old, testData,testCon)
	fmt.Println(unsafe.Sizeof(uint8(1)))
}

func init() {
	fmt.Print("初始bscic包")
}

func getStr(str string) string {
	return str + "你好"
}
