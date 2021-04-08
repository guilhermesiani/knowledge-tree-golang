package algorithms

func selectionSort(toSort []int) []int {
	var sorted []int
	for range toSort {
		minorValue := findMinorValue(toSort)
		sorted = append(sorted, toSort[minorValue])
		removeSliceElement(&toSort, minorValue)
	}
	return sorted
}

func findMinorValue(toFind []int) int {
	minorValue := toFind[0]
	minorIndex := 0

	for i := range toFind {
		if toFind[i] < minorValue {
			minorValue = toFind[i]
			minorIndex = i
		}
	}

	return minorIndex
}

func removeSliceElement(toRemove *[]int, i int) {
	(*toRemove)[len(*toRemove)-1], (*toRemove)[i] = (*toRemove)[i], (*toRemove)[len(*toRemove)-1]
	*toRemove = (*toRemove)[:len(*toRemove)-1]
}
