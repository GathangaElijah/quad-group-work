package piscine

import "fmt"

func QuadA(x,y int){
	if x < 0 || y < 0 {
		return
	}

	for i := 1; i <= y; i++{
		if i == 1 || i == y {
			for j := 0; j <= x; j++ {
				if j == 0 || j == x{
					fmt.Print("o")
				}else{
					fmt.Print("-")
				}
			}
			fmt.Println()
		}else{
			for j := 0; j <= x-1; j++ {
				if j == 0 || j == x-1 {
					fmt.Print("|")
				}else{
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} 
	}
