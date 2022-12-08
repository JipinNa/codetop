package leetcode

func Merge(nums1 []int, m int, nums2 []int, n int) {
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, len(nums1))
	for i, j := 0, 0; i < m || j < n; {
		if i == m {
			nums3[i+j] = nums2[j]
			j++
		} else if j == n {
			nums3[i+j] = nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			nums3[i+j] = nums2[j]
			j++
		} else {
			nums3[i+j] = nums1[i]
			i++
		}
	}
	for i := range nums3 {
		nums1[i] = nums3[i]
	}
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	isChanged := false
	for i, j := 0, 0; j < n; i++ {
		if isChanged {
			nextJ := j + 1
			if nextJ != n && nums2[j] > nums2[nextJ] {
				nums2[j], nums2[nextJ] = nums2[nextJ], nums2[j]
			} else {
				isChanged = false
			}
		}
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			isChanged = true
		} else {
			j++
		}
	}
}
