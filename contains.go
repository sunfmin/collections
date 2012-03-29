package collections

func StringContainsAny(cs []string, es ...string) bool {
	for _, c := range cs {
		for _, e := range es {
			if c == e {
				return true
			}
		}
	}

	return false
}

func Int64ContainsAny(cs []int64, es ...int64) bool {
	for _, c := range cs {
		for _, e := range es {
			if c == e {
				return true
			}
		}
	}

	return false
}

func Float64ContainsAny(cs []float64, es ...float64) bool {
	for _, c := range cs {
		for _, e := range es {

			if c == e {
				return true
			}
		}
	}

	return false
}
