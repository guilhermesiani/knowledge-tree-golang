package algorithms

func findHighestValue(toFind []int) int {
	var highest int = 0
	for _, v := range toFind {
		if v > highest {
			highest = v
		}
	}
	return highest
}
