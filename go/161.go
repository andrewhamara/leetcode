// Author: Andrew Hamara

// Solution to leetcode problem 161. Majority Element

func majorityElement(nums[] int) int {

    var candidate int = 0
    var count int = 0

    for _, n := range nums {
        
        if count == 0 {
            candidate = n
        }

        if n == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}
