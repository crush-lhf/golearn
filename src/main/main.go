package main

//自定义的包 必须在goroot/gopath 的src下 否则无法识别 调用使用包名即可
import (
	"fmt"
	"go-learn/src/basics"
	_ "go-learn/src/lihf"
)

func init() {
	fmt.Println("这是第一个go语言程序")
}

func main() {
	name := "李海峰"
	fmt.Println(name)
	count()
	basics.ShowData()
}

func count() {
	i := 1
	j := 2
	if i > j {
		fmt.Print(i)
	} else {
		fmt.Print(j)
	}

}
