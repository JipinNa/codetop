package leetcode

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	if l&1 == 1 {
		return float64(nums1[l/2])
	} else {
		return (float64(nums1[l/2]) + float64(nums1[l/2-1])) / 2
	}

}
