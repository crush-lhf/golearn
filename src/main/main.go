package main

import "fmt"
func init() {
	fmt.Println("这是第一个go语言程序")
}

func main() {
	name := "李海峰"
	fmt.Println(name)
	count()

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
