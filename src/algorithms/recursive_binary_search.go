package algorithms

import "fmt"

func recursiveBinarySearch(v []int, e int) int {
	minor := 0
	major := len(v) - 1

	if minor == 0 && minor == major {
		return -1
	}

	middle := int((minor + major) / 2)
	target := v[middle]

	if target == e {
		return middle
	}

	if target > e {
		fmt.Println("DO NOT PASS TARGET > E")
		return recursiveBinarySearch(removeHalfUp(v, middle), e) - middle
	}

	fmt.Println("PASS TARGET > E")
	return recursiveBinarySearch(removeHalfDown(v, middle), e) + middle
}

func removeHalfUp(toRemove []int, middle int) []int {
	return toRemove[:middle]
}

func removeHalfDown(toRemove []int, middle int) []int {
	return toRemove[middle:]
}
