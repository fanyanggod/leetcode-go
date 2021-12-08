package main

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	quickSort(nums, 0, length-1)
	ans := make([][]int, 0)
	for i := 0; i < length; i++ {
		left := i + 1
		right := length - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				if checkAns(ans, nums[i], nums[left], nums[right]) {
					ans = append(ans, []int{nums[i], nums[left], nums[right]})
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}

		}
	}
	return ans
}

func checkAns(ans [][]int, i, j, z int) bool {
	for k, _ := range ans {
		if ans[k][0] == i && ans[k][1] == j && ans[k][2] == z {
			return false
		}
	}
	return true
}

func quickSort(nums []int, left int, right int) {
	if left < right {
		partitionIndex := partition(nums, left, right)
		quickSort(nums, left, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, index)
			index += 1
		}
	}
	swap(nums, pivot, index-1)
	return index - 1
}
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
