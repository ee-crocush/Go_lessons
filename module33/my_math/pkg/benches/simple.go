package benches

func Simple(data []int, item int) int {
	for i := range data {
		if data[i] == item {
			return i
		}
	}
	return -1
}
