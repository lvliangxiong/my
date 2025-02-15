package sort

import (
	"fmt"
	"math/rand"
)

type Quick struct{}

// Sort adopts an three-way partition algorithm, which is effective for sorting
// array containing lots of duplicate elements.
//
// Relevant leetcode problem: https://leetcode.cn/problems/sort-an-array/.
func (q *Quick) Sort(nums []int) {
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

	q.Sort(nums[:lt+1])
	q.Sort(nums[gt:])
}

// HoareSort is the original version of quick sort
// by C.A.R.Hoare.
// This implementation has two main problems:
//  1. Pivot was not randomly chosen.
//  2. No optimization on array with lots of duplicated
//     elements.
func (q *Quick) HoareSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, pivot := -1, len(nums), nums[0] // You can change nums[0] to nums[rand.Intn(len(nums))].
	for {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 0 <= j < len(nums)-1
	// Element in nums[:j+1] <= nums[j+1:].

	q.HoareSort(nums[:j+1])
	q.HoareSort(nums[j+1:])
}

// Select find the k-th largest element in nums.
//
// Relevant leetcode problem: https://leetcode.cn/problems/kth-largest-element-in-an-array/.
func (q *Quick) Select(nums []int, k int) int {
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
		return q.Select(nums[:lt+1], k+lt-len(nums)+1)
	}
	if len(nums)-k >= gt {
		return q.Select(nums[gt:], k)
	}
	return pivot
}
