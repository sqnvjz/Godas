package main

import (
	"Godas/core"
	"fmt"
)

func main() {
	//df1 := core.DataFrame{}
	//a := [][]int{
	//	{1, 2, 3},
	//	{1, 2, 3},
	//	{1, 2, 3},
	//}
	//df2 := core.DataFrame{}
	//b := [][]int{
	//	{1, 2},
	//	{1, 2},
	//}
	df1 := core.DataFrame{}
	a := [][]string{
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a", "b", "c"},
	}
	df2 := core.DataFrame{}
	b := [][]string{
		{"d", "e"},
		{"d", "e"},
	}
	//df1.ItoDF(a)
	//df2.ItoDF(b)
	df1.StoDF(a)
	df2.StoDF(b)
	fmt.Println(df1)
	fmt.Println(df1.Shape())
	df1.Concat(df2)
	fmt.Println(df1)
	fmt.Println()
}
