package zapper

import "fmt"

type Err struct {
	Message string
	Params  []any
}

func NewError(msg string, params ...any) *Err {
	return &Err{
		Message: msg,
		Params:  params,
	}
}

func (e *Err) Error() string {
	return fmt.Sprintf(e.Message, e.Params...)
}
