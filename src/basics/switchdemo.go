package basics

import "fmt"

func switchdemo(letter string) {
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("元音")
	default:
		fmt.Println("普通")
	}
}
