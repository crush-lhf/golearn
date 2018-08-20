package basics

import "fmt"

func init() {
	fmt.Println("for循环 经典使用：")
	multiplicationTable()
}

func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, " ")
		}
		fmt.Println()
	}

}
