// Author: Andrew Hamara

// Solution to leetcode problem 1. Two Sum

func twoSum(nums []int, target int) []int {

    visited := make(map[int]int)

    for i:=0; i < len(nums); i++ {
        
        var cur int = target - nums[i]

        val, ok := visited[cur]

        if ok {
            return []int{i, val}
        }

        visited[nums[i]] = i
    }

    return nil
}
