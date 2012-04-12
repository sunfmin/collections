package collections

import (
	"math/rand"
	"reflect"
)

func RandomPick(source interface{}, result interface{}) {
	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		panic("result must be a pointer")
	}
	resultVal = resultVal.Elem()

	sourceVal := reflect.ValueOf(source)

	randIndex := rand.Intn(sourceVal.Len())
	val := sourceVal.Index(randIndex)
	resultVal.Set(val)
	return
}
