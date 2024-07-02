/**
* Given two integer arrays nums1 and nums2, return an array of their intersection. 
* Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.
* See: https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
 */
 package main

 type ListNode struct {
	 Val  int
	 Next *ListNode
 }

 
func intersect(nums1 []int, nums2 []int) []int {
    
    count := map[int]int {}

    for _, n := range nums1 {
        count[n]++
    }

    ans := []int{}

    for _, n := range nums2 {
        if count[n] > 0 {
            ans = append(ans, n)
            count[n]--s
            
        }
    }

    return ans
}