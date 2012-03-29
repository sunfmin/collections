package collections

import (
	"reflect"
)

func StringSelect(cs []string, filterFunc func(e string) bool) (rs []string) {
	for _, c := range cs {
		if filterFunc(c) {
			rs = append(rs, c)
		}
	}

	return
}

func Int64Select(cs []int64, filterFunc func(e int64) bool) (rs []int64) {
	for _, c := range cs {
		if filterFunc(c) {
			rs = append(rs, c)
		}
	}

	return
}

func Float64Select(cs []float64, filterFunc func(e float64) bool) (rs []float64) {
	for _, c := range cs {
		if filterFunc(c) {
			rs = append(rs, c)
		}
	}

	return
}

func Collect(source interface{}, mapFunc func(e interface{}) (interface{}, bool), result interface{}) {
	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		panic("result must be a pointer")
	}
	resultVal = resultVal.Elem()

	sourceVal := reflect.ValueOf(source)

	newslice := reflect.MakeSlice(resultVal.Type(), 0, 0)
	for i := 0; i < sourceVal.Len(); i++ {
		val := sourceVal.Index(i)
		converted, add := mapFunc(val.Interface())
		convertedVal := reflect.ValueOf(converted)
		if add {
			newslice = reflect.Append(newslice, convertedVal)
		}
	}
	resultVal.Set(newslice)
	return
}
