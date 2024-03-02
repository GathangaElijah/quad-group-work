package piscine

import "fmt"

func QuadA_3(x,y int){
	if x < 0 || y < 0 {
		return
	}

	for i := 1; i <= y; {
		fmt.Print("o")
	}
}