package basics

import (
	"fmt"
	"unsafe"
)

//方法名不大写的话 其他包无法识别
func ShowData() {
	fmt.Println(unsafe.Sizeof(uint8(1)))
}
