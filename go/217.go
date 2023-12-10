// Author: Andrew Hamara

// Solution to leetcode problem 217. Contains Duplicate

func containsDuplicate(nums []int) bool {
    
    mySet := make(map[int]bool)

    for _, value := range nums {
        mySet[value] = true
    }

    return len(mySet) != len(nums)
}
