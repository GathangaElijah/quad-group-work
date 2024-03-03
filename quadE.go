package piscine

import "fmt"

func QuadE(x, y int) {
	if x <= 0 && y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				//first corner
				fmt.Print("A")
			} else if i == 1 && j == x {
				//second corner
				fmt.Print("C")
			} else if i == y && j == 1 {
				// third corner
				fmt.Print("C")
			} else if i == y && j == x {
				//fourth corner
				fmt.Print("A")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
