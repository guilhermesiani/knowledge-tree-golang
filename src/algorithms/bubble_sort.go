package algorithms

func bubbleSort(toSort [5]int) [5]int {
	sorted := toSort
	for i := 0; i < len(toSort); i++ {
		toSort := sorted
		for j := 0; j < len(toSort); j++ {
			if (len(toSort)-(i+1)) >= j && sorted[j] > sorted[len(toSort)-(i+1)] {
				toSort = sorted;
				sorted[len(toSort)-(i+1)] = sorted[j]
				sorted[j] = toSort[len(toSort)-(i+1)]
			}
		}
	}

	return sorted
}