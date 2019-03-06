package core

import "reflect"

type GoPipe struct {
	Data  []interface{}
	GType string
}

type GoFrame struct {
	Data    []GoPipe
	TypeMap map[int]string
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
	return nil
}

func (gp *GoPipe) Init(ele []interface{}) error {
	for _, v := range ele {
		gp.Data = append(gp.Data, v)
		gp.GType = reflect.TypeOf(v).String()
	}
	return nil
}
