package sort

import (
	"fmt"
	"math/rand"
)

type Quick struct{}

// Sort adopts an tree-way partition algorithm, which is effective for sorting
// array containing lots of duplicate elements.
//
// Relevant leetcode problem: https://leetcode.cn/problems/sort-an-array/description/.
func (quick *Quick) Sort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	randIdx := rand.Intn(len(nums))
	nums[0], nums[randIdx] = nums[randIdx], nums[0]

	// Element in (-inf, lt] < pivot.
	// Element in (lt, i) == pivot.
	// Element in [i, gt) is undetermined.
	// Element in [gt, +inf) > pivot.
	lt, gt, i, pivot := -1, len(nums), 1, nums[0]
	for i < gt {
		if nums[i] == pivot {
			i++
		} else if nums[i] < pivot {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		}
	}

	quick.Sort(nums[0 : lt+1])
	quick.Sort(nums[gt:])
}

func (quick *Quick) TwoWaySort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	randIdx := rand.Intn(len(nums))
	nums[0], nums[randIdx] = nums[randIdx], nums[0]

	// Ele in range [..i] <= pivot.
	// Ele in range [j..] >= pivot.
	// Ele in range (i..j) is undetermined.
	//
	// In the last iter, we got i >= j and we don't swap them.
	// And there are only two cases:
	//   1. i == j and nums[i] == nums[j] == pivot
	//   2. j == i-1 and nums[j] <= pivot and nums[i] >= pivot
	i, j, pivot := 0, len(nums), nums[0]
	for i < j {
		for i++; i < j && nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[0], nums[j] = nums[j], nums[0]
	// nums[j] is the partition location.
	quick.TwoWaySort(nums[0:j])
	quick.TwoWaySort(nums[j+1:])
}

// Select find the k-th largest element in nums.
//
// Relevant leetcode problem: https://leetcode.cn/problems/kth-largest-element-in-an-array/.
func (quick *Quick) Select(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		panic(fmt.Sprintf("unexpected k = %d", k))
	}

	if len(nums) == 1 {
		return nums[0]
	}

	randIdx := rand.Intn(len(nums))
	nums[0], nums[randIdx] = nums[randIdx], nums[0]

	lt, gt, i, pivot := -1, len(nums), 1, nums[0]
	for i < gt {
		if nums[i] == pivot {
			i++
		} else if nums[i] < pivot {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		}
	}

	// Element in range (le, gt) equals to pivot.
	if len(nums)-k <= lt {
		return quick.Select(nums[0:lt+1], k+lt-len(nums)+1)
	}
	if len(nums)-k >= gt {
		return quick.Select(nums[gt:], k)
	}
	return pivot
}
