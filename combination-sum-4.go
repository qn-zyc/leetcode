package leetcode

import "sort"

/*
<https://leetcode.com/problems/combination-sum-iv/>

Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.

Example:

nums = [1, 2, 3]
target = 4

The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

Note that different sequences are counted as different combinations.

Therefore the output is 7.
*/

// FAILED 超时

func CombinationSum4(nums []int, target int) int {
	sort.Ints(nums)

	return combinationSum4Helper(nums, target)
}

func combinationSum4Helper(nums []int, target int) int {
	c := 0
	if target <= 0 {
		return 0
	}
	for _, n := range nums {
		if n > target {
			break
			//continue
		}
		if n == target {
			c++
			break
			//continue
		}
		c += combinationSum4Helper(nums, target-n)
	}
	return c
}

// https://discuss.leetcode.com/topic/52186/my-3ms-java-dp-solution

func CombinationSum4_2(nums []int, target int) int {
	sort.Ints(nums)

	temp := make([]int, target+1)
	for i := 0; i < len(temp); i++ {
		for _, n := range nums {
			if n > i {
				break
			}
			if n == i {
				temp[i]++
			} else {
				temp[i] += temp[i-n]
			}
		}
	}
	return temp[target]
}
