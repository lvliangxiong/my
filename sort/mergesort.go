package sort

type Merge struct{}

func (m *Merge) Sort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	start, end := 0, len(nums)
	mid := (end-start)>>1 + start
	m.Sort(nums[start:mid])
	m.Sort(nums[mid:end])
	m.merge(nums, mid)
}

// merge two ordered nums[0:c] & nums[c:]
func (m *Merge) merge(nums []int, c int) {
	// [0, c) and [c, len(nums))
	if c <= 0 || c >= len(nums) {
		return
	}
	// Actually you can make a smaller slice for temporary storage of nums, but
	// it will make code much more complicated, hard to read. Here we just
	// provide a simpler solution, which consumes more space although.
	tmp := make([]int, len(nums))
	p1, p2, cur := c-1, len(nums)-1, len(nums)-1
	for p1 >= 0 && p2 >= c {
		if nums[p1] < nums[p2] {
			tmp[cur] = nums[p2]
			p2--
		} else {
			tmp[cur] = nums[p1]
			p1--
		}
		cur--
	}
	for p1 >= 0 {
		tmp[cur] = nums[p1]
		p1--
		cur--
	}
	for p2 >= c {
		tmp[cur] = nums[p2]
		p2--
		cur--
	}
	copy(nums, tmp)
}
