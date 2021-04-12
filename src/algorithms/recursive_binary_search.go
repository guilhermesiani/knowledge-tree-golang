package algorithms

func recursiveBinarySearch(v []int, e int) int {
	var result int
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
		result = recursiveBinarySearch(removeHalfUp(v, middle), e)
	} else {
		result = recursiveBinarySearch(removeHalfDown(v, middle), e)
		if result != -1 {
			return result + middle
		}
	}

	return result
}

func removeHalfUp(toRemove []int, middle int) []int {
	return toRemove[:middle]
}

func removeHalfDown(toRemove []int, middle int) []int {
	return toRemove[middle:]
}
