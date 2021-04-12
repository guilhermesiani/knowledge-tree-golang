package algorithms

func recursiveBinarySearch(v []int, e int) int {
	minor := 0
	major := len(v) - 1

	if minor == major {
		return -1
	}

	middle := int((minor + major) / 2)
	target := v[middle]

	if target == e {
		return middle
	}

	if target > e {
		return recursiveBinarySearch(removeHalfUp(v, middle), e) - middle
	}

	return recursiveBinarySearch(removeHalfDown(v, middle), e) + middle
}

func removeHalfUp(toRemove []int, middle int) []int {
	return toRemove[:middle]
}

func removeHalfDown(toRemove []int, middle int) []int {
	return toRemove[middle:]
}
