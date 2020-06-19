package algorithms

func bubbleSort(toSort [5]int) [5]int {
	sorted := toSort
	for i := 0; i < len(toSort); i++ {
		toSort := sorted
		for j := 0; j < len(toSort); j++ {
			pos := len(toSort)-(i+1)
			if pos >= j && sorted[j] > sorted[pos] {
				toSort = sorted;
				sorted[pos] = sorted[j]
				sorted[j] = toSort[pos]
			}
		}
	}

	return sorted
}