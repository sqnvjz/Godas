package main

import (
	"GoDas/core"
	"fmt"
)

func main() {
	gp1 := core.NewGoPipe()
	t := []int{1, 2, 3, 4, 5}
	a := make([]interface{}, len(t))
	for i, v := range t {
		a[i] = v
	}
	gp1.Init(a)
	fmt.Println(gp1)

	gp2 := core.NewGoPipe()
	f := []string{"a", "b", "c", "d", "ef"}
	a = make([]interface{}, len(f))
	for i, v := range f {
		a[i] = v
	}
	gp2.Init(a)
	fmt.Println(gp2)

	df := core.NewGoFrame()
	df.Add(*gp1)
	df.Add(*gp2)
	fmt.Println(df)

}
