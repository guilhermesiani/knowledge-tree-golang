package algorithms

func binarySearch(x int, v [4]int, e int, d int) int {
	var middle int
	middle = (e + d) / 2
	if v[middle] == x {
		return middle
	}
	if e >= d {
		return -1 // not found
	} else if v[middle] < x {
		return binarySearch(x, v, middle+1, d)
	} else {
		return binarySearch(x, v, e, middle-1)
	}
}