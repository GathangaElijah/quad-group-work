package piscine

import "fmt"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				//First corner
				fmt.Print("A")
			} else if i == 1 && j == x {
				//Second corner
				fmt.Print("C")
			} else if i == y && j == 1 {
				// Third corner
				fmt.Print("A")
			} else if i == y && j == x {
				//Fourth corner
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
