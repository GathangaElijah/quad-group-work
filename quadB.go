package piscine

import "fmt"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			// For the first corner
			if i == 1 && j == 1 {
				fmt.Print("/")
				//For the second corner
			} else if i == 1 && j == x {
				fmt.Print("\\")
				//For the third corner
			} else if i == y && j == 1 {
				fmt.Print("\\")
				//For the fourth corner
			} else if i == y && j == x {
				fmt.Print("/")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
