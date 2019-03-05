package core

import (
	"errors"
	"Godas/utils"
)

type DataFrame struct {
	Data    interface{}
	Index   int
	Columns int
	Dtype   string
}

func (df *DataFrame) ItoDF(arr [][]int) error {
	df.Data = arr
	rows := 0
	cols := 0
	for _, r := range arr {
		rows++
		if rows == 1 {
			for _, c := range r {
				_ = c
				cols++
			}
		}
	}
	df.Index = rows
	df.Columns = cols
	df.Dtype = "int"
	return nil
}

func (df *DataFrame) StoDF(arr [][]string) error {
	df.Data = arr
	rows := 0
	cols := 0
	for _, r := range arr {
		rows++
		if rows == 1 {
			for _, c := range r {
				_ = c
				cols++
			}
		}
	}
	df.Index = rows
	df.Columns = cols
	df.Dtype = "string"
	return nil
}

type Shape struct {
	Length int
	Width  int
}

func (df *DataFrame) Shape() *Shape {
	var s Shape
	s.Length = df.Index
	s.Width = df.Columns
	return &s
}

func (df *DataFrame) Concat(comp DataFrame) error {
	if df.Dtype != comp.Dtype {
		return errors.New("Incompatible Dtype.")
	}

	switch df.Dtype {
	case "int":
		a := df.Data.([][]int)
		b := comp.Data.([][]int)
		var target [][]int
		newrow := utils.Max(len(a), len(b))
		newcol := len(a[0]) + len(b[0])
		minrow := utils.Min(len(a), len(b))
		for i := 0; i < newrow; i++ {
			var m []int
			if i < minrow {
				for j := 0; j < newcol; j++ {
					if j < len(a) {
						m = append(m, a[i][j])
					} else {
						m = append(m, b[i][j-len(a)])
					}
				}
			} else {
				if len(a) == minrow {
					for j := 0; j < newcol; j++ {
						if j < len(a) {
							m = append(m, 0)
						} else {
							m = append(m, b[i][j-len(a)])
						}
					}
				} else {
					for j := 0; j < newcol; j++ {
						if j < len(a) {
							m = append(m, a[i][j])
						} else {
							m = append(m, 0)
						}
					}
				}
			}

			target = append(target, m)
		}
		df.Data = target
		df.Index = newrow
		df.Columns = newcol
	case "string":
		a := df.Data.([][]string)
		b := comp.Data.([][]string)
		var target [][]string
		newrow := utils.Max(len(a), len(b))
		newcol := len(a[0]) + len(b[0])
		minrow := utils.Min(len(a), len(b))
		for i := 0; i < newrow; i++ {
			var m []string
			if i < minrow {
				for j := 0; j < newcol; j++ {
					if j < len(a) {
						m = append(m, a[i][j])
					} else {
						m = append(m, b[i][j-len(a)])
					}
				}
			} else {
				if len(a) == minrow {
					for j := 0; j < newcol; j++ {
						if j < len(a) {
							m = append(m, "")
						} else {
							m = append(m, b[i][j-len(a)])
						}
					}
				} else {
					for j := 0; j < newcol; j++ {
						if j < len(a) {
							m = append(m, a[i][j])
						} else {
							m = append(m, "")
						}
					}
				}
			}

			target = append(target, m)
		}
		df.Data = target
		df.Index = newrow
		df.Columns = newcol
	}
	return nil
}
