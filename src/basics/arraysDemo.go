package basics

import "fmt"

func init() {
	a := [3][2]string{
		{"lihf", "fufu"},
		{"ruqin", "ff"},
		{"feilong", "wl"},
	}
	printArray(a)
	printUnCheck()

}
func printArray(a [3][2]string) {
	for _, v := range a {
		for j, v2 := range v {
			fmt.Println("%d value %.2f", j, v2)
		}
	}
}
func printUnCheck() {
	b := [6]int{1, 2, 2, 3, 4, 5,}
	for i, v := range b {
		fmt.Println(i, v)
	}
}
