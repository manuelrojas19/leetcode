/**
 * Given an array of characters chars, compress it using the following algorithm:
 * Begin with an empty string s. For each group of consecutive repeating characters in chars:
 * - If the group's length is 1, append the character to s.
 * - Otherwise, append the character followed by the group's length.
 * The compressed string s should not be returned separately, but instead, be stored in the input character array chars.
 * Note that group lengths that are 10 or longer will be split into multiple characters in chars.
 * After modifying the input array chars, the function returns the new length of the array.
 * You must write an algorithm that uses only constant extra space.
 * See: https://leetcode.com/problems/string-compression/description/
 */
package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	// Initialize variables
	right := 0 // Pointer to current position in result array
	count := 1 // Count of consecutive characters

	// Iterate through the array
	for left := 0; left < len(chars); left++ {
		// Check if current character is different from previous one or end of array
		if left == len(chars)-1 || chars[left] != chars[left+1] {
			// Assign current character to result array
			chars[right] = chars[left]
			right++

			// Append count if greater than 1
			if count > 1 {
				countStr := strconv.Itoa(count) // Convert count to string
				for _, ch := range countStr {
					chars[right] = byte(ch) // Append each character of count string
					right++
				}
			}
			// Reset count for new character group
			count = 1
		} else {
			// Increment count for consecutive characters
			count++
		}
	}

	// Return length of compressed array
	return right
}

func main() {
	// Example usage
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'c'}
	newLength := compress(chars)
	fmt.Println(string(chars[:newLength])) // Output should be "ab10c"
	fmt.Println(newLength)                 // Output should be the new length
}
