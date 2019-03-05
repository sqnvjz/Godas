package core

import (
	"errors"
	"Godas/utils"
)

type DataFrame struct {
	Data      interface{}
	Length    int
	Width     int
	Dtype     string
	Indices   []string
	Columns   []string
	IndexMap  map[string]int
	ColumnMap map[string]int
}

func NewDataFrame() *DataFrame {
	return &DataFrame{}
}

func (df *DataFrame) InitIndices(index []string) error {
	df.Indices = index
	m := make(map[string]int)
	for i, s := range index {
		m[s] = i
	}
	df.IndexMap = m
	return nil
}

func (df *DataFrame) InitColumns(column []string) error {
	df.Columns = column
	m := make(map[string]int)
	for i, s := range column {
		m[s] = i
	}
	df.ColumnMap = m
	return nil
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
	df.Length = rows
	df.Width = cols
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
	df.Length = rows
	df.Width = cols
	df.Dtype = "string"
	return nil
}

type Shape struct {
	Length int
	Width  int
}

func (df *DataFrame) Shape() *Shape {
	var s Shape
	s.Length = df.Length
	s.Width = df.Width
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
		df.Length = newrow
		df.Width = newcol
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
		df.Length = newrow
		df.Width = newcol
	}
	return nil
}

func (df *DataFrame) Transpose() error {
	switch df.Dtype {
	case "int":
		arr := df.Data.([][]int)
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr[i]); j++ {
				if i != j {
					temp := arr[i][j]
					arr[i][j] = arr[j][i]
					arr[j][i] = temp
				}
			}
		}
		df.Data = arr
	case "string":
		arr := df.Data.([][]string)
		for i := 0; i < len(arr); i++ {
			for j := i; j < len(arr[i]); j++ {
				if i != j {
					temp := arr[i][j]
					arr[i][j] = arr[j][i]
					arr[j][i] = temp
				}
			}
		}
		df.Data = arr
	}
	return nil
}

func (df *DataFrame) Cut(rows []int, cols []int) error {
	switch df.Dtype {
	case "int":
		var cut [][]int
		arr := df.Data.([][]int)
		for i := 0; i < len(arr); i++ {
			var m []int
			for j := 0; j < len(arr[0]); j++ {
				if utils.InArray(i, rows) && utils.InArray(j, cols) {
					m = append(m, arr[i][j])
				}
			}
			if len(m) > 0 {
				cut = append(cut, m)
			}
		}
		df.Data = cut
		df.Length = len(rows)
		df.Width = len(cols)
	case "string":
		var cut [][]string
		arr := df.Data.([][]string)
		for i := 0; i < len(arr); i++ {
			var m []string
			for j := 0; j < len(arr[0]); j++ {
				if utils.InArray(i, rows) && utils.InArray(j, cols) {
					m = append(m, arr[i][j])
				}
			}
			if len(m) > 0 {
				cut = append(cut, m)
			}
		}
		df.Data = cut
		df.Length = len(rows)
		df.Width = len(cols)
	}
	return nil
}
