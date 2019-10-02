package util

func SwapInt(slice []int, i1 int, i2 int) {
	slice[i1], slice[i2] = slice[i2], slice[i1]
}

func ShiftRightAndPutFirstInt(slice []int, index int) {
	x := slice[index]
	for i := index; i > 0; i-- {
		slice[i] = slice[i-1]
	}
	slice[0] = x
}
