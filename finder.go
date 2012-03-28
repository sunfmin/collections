package collections

import (
	"reflect"
)

func Find(source interface{}, finderFunc func(e interface{}) bool, result interface{}) {
	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		panic("result must be a pointer")
	}
	resultVal = resultVal.Elem()

	sourceVal := reflect.ValueOf(source)
	for i := 0; i < sourceVal.Len(); i++ {
		val := sourceVal.Index(i)
		if finderFunc(val.Interface()) {
			resultVal.Set(val)
			return
		}
	}
}
