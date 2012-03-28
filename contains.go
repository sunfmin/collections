package collections

func StringsContainAny(cs []string, es ...string) bool {
	for _, c := range cs {
		for _, e := range es {
			if c == e {
				return true
			}
		}
	}

	return false
}

func Int64sContainAny(cs []int64, es ...int64) bool {
	for _, c := range cs {
		for _, e := range es {
			if c == e {
				return true
			}
		}
	}

	return false
}

func Float64sContainAny(cs []float64, es ...float64) bool {
	for _, c := range cs {
		for _, e := range es {

			if c == e {
				return true
			}
		}
	}

	return false
}
