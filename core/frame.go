package core

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
