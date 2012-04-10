package collections

func StringInGroupsOf(cs []string, size int) (r [][]string) {
	ri := make([]string, 0)
	for i, c := range cs {
		if i > 0 && i%size == 0 {
			r = append(r, ri)
			ri = make([]string, 0)
		}
		ri = append(ri, c)
	}
	r = append(r, ri)
	return
}
