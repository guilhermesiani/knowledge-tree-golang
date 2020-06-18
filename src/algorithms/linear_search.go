package algorithms

func linearSearch(vector [5]string, toSearch string) int {
	for i := 0; i < len(vector); i++ {
		if toSearch == vector[i] {
			return i
		}
	}
	return -1
}