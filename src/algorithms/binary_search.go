package algorithms

import (
	"fmt"
	"math"
)

func binarySearch(v []int, e int) int {
	minor := 0
	major := len(v) - 1

	fmt.Println(minor)
	fmt.Println(major)

	for minor <= major {
		middle := int(math.Ceil(float64(minor) + float64(major)/float64(2)))
		target := v[middle]

		fmt.Println(middle)
		fmt.Println(target)
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
