/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}

		}
	}
	return nil
}

/*
beats 100%
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)//空のmap生成
	for i, n := range nums {//i:回数,key/n:value
			_, prs := m[n]//key:"n"のvalueにアクセス
			if prs {
					return []int{m[n], i}
			} else {
					m[target-n] = i
			}
	}
	return nil
}

simple solution
func twoSum(nums []int, target int) []int {
	conjugateSum := make(map[int]int)

	for i, num := range nums {
			if j, ok := conjugateSum[num]; ok {
					return []int{j, i}
			}
			conjugateSum[target - num] = i
	}
	return nil
}
*/

