package main

import "fmt"

// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a substring

func lengthOfLongestSubstring(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i]==s[i+1] {
			i=1
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && i == j {
				continue
			} else {
				result = s[i:j]
			}
		}
	}
	return result
}

func main3() {
	name := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(name))
}
