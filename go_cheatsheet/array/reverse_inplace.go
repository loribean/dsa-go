package array

// ReverseInPlace solves the problem in O(n) time and O(1) space.
func ReverseInPlace(list []int, start, end int) {
	// to solve this question, we need to keep swapping start and end indexes till we reach the middle
	//for the stop condition, we want to stop when we reach the middle
	//how do we know we have reached the middle? when i > j
	//what is j? j is the mirror index of i
	// j := start+end-i
	// NOTE: j cannot be end -i, this only works if start is 0.
	// i < j is the continue
	for i := start; i < start+end-i; i++ {
		j := start + end - i
		list[i], list[j] = list[j], list[i]
	}

	//simpler, two pointers
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

}
