package main

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	for i, v := range nums {
		if value, ok := valueMap[target-v]; ok {
			return []int{value, i}
		}
		valueMap[nums[i]] = i
	}
	return nil
}
