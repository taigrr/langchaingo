package sliceutil

// MinInt returns the minimum value in nums.
// If nums is empty, it returns 0.
func MinInt(nums []int) int {
	var m int
	for idx := range nums {
		item := nums[idx]
		if idx == 0 {
			m = item
			continue
		}
		if item < m {
			m = item
		}
	}
	return m
}
