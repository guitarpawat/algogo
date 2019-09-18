package sorted_search

func BinarySearch(target int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}

	mid := len(numbers)/2
	if target == numbers[mid] {
		return true
	} else if target > numbers[mid] {
		return BinarySearch(target, numbers[mid+1:])
	} else {
		return BinarySearch(target, numbers[:mid])
	}
}