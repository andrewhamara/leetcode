// Author: Andrew Hamara

// Solution to leetcode problem 125. Valid Palindrome

import (
    "regexp"
)

func isPalindrome(s string) bool {
    reg := regexp.MustCompile("[^A-Za-z0-9]+")
    alnum := strings.ToLower(reg.ReplaceAllString(s, ""))

    left := 0
    right := len(alnum) - 1

    for right > left {
        if alnum[left] != alnum[right] {
            return false
        }
        left += 1
        right -= 1
    }
    return true

}
