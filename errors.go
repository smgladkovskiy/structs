package structs

import (
	"fmt"
)

const typeIsNotAcceptable = "unsupported data type `%T` for .Scan(), storing driver.Value into type %T"

type TypeIsNotAcceptable struct {
	CheckedValue interface{}
	CheckedType  interface{}
}

func (err TypeIsNotAcceptable) Error() string {
	return fmt.Sprintf(typeIsNotAcceptable, err.CheckedValue, err.CheckedType)
}

const valueIsNotAcceptable = "unsupported data value `%+v` for .Scan(), storing driver.Value into type %T"

type ValueIsNotAcceptable struct {
	CheckedValue interface{}
	CheckedType  interface{}
}

func (err ValueIsNotAcceptable) Error() string {
	return fmt.Sprintf(valueIsNotAcceptable, err.CheckedValue, err.CheckedType)
}
