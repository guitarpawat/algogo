package sorted_search

func LinearSearch(target int, numbers []int) bool {
	for _, v := range numbers {
		if v == target {
			return true
		}
	}
	return false
}
