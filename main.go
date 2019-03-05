package main

import (
	"Godas/core"
	"fmt"
)

func main() {
	df1 := core.DataFrame{}
	a := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	df1.ItoDF(a)
	fmt.Println(df1)
	fmt.Println(df1.Shape())
}
