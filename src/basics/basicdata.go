package basics

import (
	"fmt"
	"unsafe"
)

func showData()  {
	fmt.Println(unsafe.Sizeof(uint8(1)))
}
