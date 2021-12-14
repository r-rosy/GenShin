package test

import (
	"fmt"
	"reflect"
)

func ShowType(val interface{}) {
	Type := reflect.TypeOf(val)
	fmt.Println(Type)
}
