func rob(nums []int) int {
    var lastRobbed int = 0
    var twoAgo int = 0
    var curRob int = -1000000

    for _, n := range nums {
        if n + twoAgo > lastRobbed {
            curRob = n + twoAgo
        }else {
            curRob = lastRobbed
        }
        twoAgo = lastRobbed
        lastRobbed = curRob
    }
    return curRob
}
