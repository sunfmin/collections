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

func Shuffle(source interface{}, result interface{}) {
	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		panic("result must be a pointer")
	}
	resultVal = resultVal.Elem()

	sourceVal := reflect.ValueOf(source)

	newslice := reflect.MakeSlice(resultVal.Type(), 0, 0)

	var indexes []int
	for i := 0; i < sourceVal.Len(); i++ {
		indexes = append(indexes, i)
	}
	shuffleIndexes := IntShuffle(indexes)
	for _, i := range shuffleIndexes {
		val := sourceVal.Index(i)
		newslice = reflect.Append(newslice, val)
	}

	resultVal.Set(newslice)
	return
}

func IntShuffle(source []int) (result []int) {
	shinks := source

	for {
		left := len(shinks)
		if left == 0 {
			return
		}
		selected := rand.Intn(left)
		excepted := []int{}
		for i, v := range shinks {
			if i != selected {
				excepted = append(excepted, v)
			} else {
				result = append(result, v)
			}
		}
		shinks = excepted

	}
	return
}
