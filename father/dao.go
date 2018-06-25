package father


import "reflect"

type TableName interface {
	TableName() string
}

type Daoer interface {
	ColIdx() []int
	Before() interface{}
}

type Dao struct {
	colIdx   []int
	original *reflect.Value
}

func (d *Dao) ColIdx() []int {
	return d.colIdx
}

func (d *Dao) Before() interface{} {
	if d.original == nil {
		return nil
	}

	return d.original.Elem().Addr().Interface()
}
