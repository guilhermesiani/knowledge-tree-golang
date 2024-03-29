package algorithms

func binarySearch(v []int, e int) int {
	minor := 0
	major := len(v) - 1

	for minor <= major {
		middle := int((minor + major) / 2)
		target := v[middle]

		if target == e {
			return middle
		}

		if target > e {
			major = middle - 1
		} else {
			minor = middle + 1
		}
	}

	return -1
}
