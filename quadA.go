package piscine

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			for j := 0; j < x; j++ {
				if j == 0 || j == x-1 {
					fmt.Print("o")
				} else {
					fmt.Print("-")
				}
			}
		} else {
			for j := 0; j < x; j++ {
				if j == 0 || j == x-1 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
