package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 && n == 0 {
		return 0
	}
	if (m+n)%2 == 0 {
		left := (m + n + 1) / 2
		right := (m + n + 2) / 2
		return float64(findKey(nums1, 0, m-1, nums2, 0, n-1, left)+findKey(nums1, 0, m-1, nums2, 0, n-1, right)) / 2
	} else {
		return float64(findKey(nums1, 0, m-1, nums2, 0, n-1, (m+n+1)/2))
	}
}

func findKey(nums1 []int, n1l int, n1h int, nums2 []int, n2l int, n2h int, ki int) int {
	m := n1h - n1l + 1
	n := n2h - n2l + 1
	if m > n {
		return findKey(nums2, n2l, n2h, nums1, n1l, n1h, ki)
	}
	if m == 0 {
		return nums2[n2l+ki-1]
	}
	if ki == 1 {
		if nums1[n1l] <= nums2[n2l] {
			return nums1[n1l]
		} else {
			return nums2[n2l]
		}
	}
	i := n1l - 1
	j := n2l - 1
	if m < ki/2 {
		i += m
	} else {
		i += ki / 2
	}
	if n < ki/2 {
		j += n
	} else {
		j += ki / 2
	}

	if nums1[i] > nums2[j] {
		return findKey(nums1, n1l, n1h, nums2, j+1, n2h, ki-(j-n2l+1))
	} else {
		return findKey(nums1, i+1, n1h, nums2, n2l, n2h, ki-(i-n1l+1))
	}
}
