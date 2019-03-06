package core

import (
	"GoDas/utils"
	"fmt"
	"reflect"
	"strconv"
)

type GoPipe struct {
	Data  []interface{}
	GType string
}

type GoFrame struct {
	Data      []GoPipe
	TypeMap   map[int]string
	Length    int
	Width     int
	Indices   []string
	Columns   []string
	IndexMap  map[string]int
	ColumnMap map[string]int
}

func NewGoPipe() *GoPipe {
	return &GoPipe{}
}

func NewGoFrame() *GoFrame {
	return &GoFrame{}
}

func (gf *GoFrame) Add(gp GoPipe) error {
	if gf.TypeMap == nil {
		gf.TypeMap = make(map[int]string)
	}
	gf.TypeMap[len(gf.Data)] = gp.GType
	gf.Data = append(gf.Data, gp)
	gf.Length = utils.Max(gf.Length, len(gp.Data))
	gf.Width++
	return nil
}

func (gf *GoFrame) AttachColumns(arr []string) error {
	gf.Columns = arr
	if gf.ColumnMap == nil {
		gf.ColumnMap = make(map[string]int)
	}
	for i := 0; i < gf.Width; i++ {
		gf.ColumnMap[arr[i]] = i
	}
	return nil
}

func (gf *GoFrame) AttachIndices(arr []string) error {
	gf.Indices = arr
	if gf.IndexMap == nil {
		gf.IndexMap = make(map[string]int)
	}
	for i := 0; i < gf.Length; i++ {
		gf.IndexMap[arr[i]] = i
	}
	return nil
}

func (gf *GoFrame) V() error {
	var arr [][]string
	for i := 0; i < gf.Length; i++ {
		m := make([]string, 0)
		for j := 0; j < gf.Width; j++ {
			m = append(m, "Nil")
		}
		arr = append(arr, m)
	}

	j := 0
	for _, gp := range gf.Data {
		i := 0
		for _, item := range gp.Data {
			switch gp.GType {
			case "int":
				arr[i][j] = strconv.Itoa(item.(int))
			case "string":
				arr[i][j] = item.(string)
			case "bool":
				if item == true {
					arr[i][j] = "true"
				} else {
					arr[i][j] = "false"
				}
			}
			i++
		}
		j++
	}

	for i := 0; i < gf.Length; i++ {
		for j := 0; j < gf.Width; j++ {
			fmt.Printf("%s\t\t", arr[i][j])
		}
		fmt.Println()
	}
	return nil
}

func (gp *GoPipe) Init(ele []interface{}) error {
	for _, v := range ele {
		gp.Data = append(gp.Data, v)
		gp.GType = reflect.TypeOf(v).String()
	}
	return nil
}
