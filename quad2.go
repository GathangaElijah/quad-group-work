package piscine

import "fmt"

func QuadA_2(x,y int){
	if x < 0 || y < 0 {
		return
	}
	for i := 0; i <= y; i++ {
		if i == 0 || i == y {
			for j := 0; j <= x-1; j++ {
				if j == 0 || j == x-1{
					fmt.Print("o")
				}else{
					fmt.Print("-")
				}
			}
		}
	}
}