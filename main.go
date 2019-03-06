package main

import (
	"GoDas/core"
	"fmt"
)

func main() {
	gp1 := core.NewGoPipe()
	t1 := []int{1, 2, 3, 4, 5}
	a := make([]interface{}, len(t1))
	for i, v := range t1 {
		a[i] = v
	}
	gp1.Init(a)
	fmt.Println(gp1)

	gp2 := core.NewGoPipe()
	t2 := []string{"a", "b", "c", "d", "e", "f"}
	a = make([]interface{}, len(t2))
	for i, v := range t2 {
		a[i] = v
	}
	gp2.Init(a)
	fmt.Println(gp2)

	gp3 := core.NewGoPipe()
	t3 := []bool{true, false, false, true}
	a = make([]interface{}, len(t3))
	for i, v := range t3 {
		a[i] = v
	}
	gp3.Init(a)
	fmt.Println(gp3)

	df := core.NewGoFrame()
	df.Add(*gp1)
	df.Add(*gp2)
	df.Add(*gp3)
	df.AttachColumns([]string{"cols1", "cols2", "cols3"})
	df.AttachIndices([]string{"index1", "index2", "index3", "index4", "index5", "index6"})
	df.V()
	fmt.Println(df)

}
