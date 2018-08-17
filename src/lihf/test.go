package lihf

import "fmt"

func init() {
	fmt.Println("初始化lihf包")
}

func CountArea(len, width float32) float32 {
	area := len * width
	return area
}
