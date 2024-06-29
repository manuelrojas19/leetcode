/**
* Given an integer array nums sorted in non-decreasing order
* Remove the duplicates in-place such that each unique element appears only once.
* The relative order of the elements should be kept the same.
* Then return the number of unique elements in nums.
* See: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
 */
package main

func removeDuplicates(nums []int) int {
	lastUnique := 0
	for current := 1; current < len(nums); current++ {
		if nums[lastUnique] != nums[current] {
			lastUnique++
			nums[lastUnique] = nums[current]
		}
	}
	return lastUnique + 1
}
