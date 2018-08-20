package basics

import "fmt"

func init() {
	test01(10)
}

//注意点 else 必须跟在if的的花括号后面 可以在if 上做初始化操作
func test01(num int) {
	if num%2 == 0 {
		fmt.Print("the number is even")
	} else {
		fmt.Println("odd")
	}

}
