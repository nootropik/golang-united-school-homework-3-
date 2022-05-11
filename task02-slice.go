package homework

import "sort"

func reverse(input []int64) (result []int64) {
	//Place your code here
	in2:= [...]int64
	n1:=copy(in2, input [0:])
	result := sort.Reverse(sort.IntSlice(n1))
	return result
}
