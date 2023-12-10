import (
    "fmt"
)
func findErrorNums(nums []int) []int {
    m := make(map[int]int)
    for _,num := range nums {
        _,ok := m[num]
        if ok {
            m[num] += 1
        }else {
            m[num] = 1
        }
    }
    rangeMax := len(nums) + 1

    dup := 0
    missing :=0
    for i:=1; i < rangeMax; i++ {
        if m[i] == 2 {
            dup = i
        }else if m[i] == 0 {
            missing = i
        }
    }
    return []int{dup,missing}
}
