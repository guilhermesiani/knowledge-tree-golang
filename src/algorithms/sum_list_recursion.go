package algorithms

func sumListRecursion(list *[]int) int {
	var sum int
	if len(*list) == 0 {
		return 0
	}
	sum = (*list)[0]
	removeSliceElement(list, 0)
	return sum + sumListRecursion(list)
}
