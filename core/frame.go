package core

import "errors"

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

func (df *DataFrame) Concat(comp *DataFrame) error {
	if df.Dtype != comp.Dtype {
		return errors.New("Incompatible Dtype.")
	}
	switch df.Dtype {
	case "int":
		var target [][]int
		df.Data = target
	case "string":
		var target [][]int
		df.Data = target
	}
	return nil
}
