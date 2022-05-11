package homework

import "sort"

func reverse(input []int64) (result []int64) {
	//Place your code here
	in2 := copy(input)
	result := sort.Reverse(sort.IntSlice(in2))
	return result
}
