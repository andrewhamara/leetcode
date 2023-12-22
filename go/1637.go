func maxWidthOfVerticalArea(points [][]int) int {
    xVals := make([]int, 0)
    for _,val := range points {
        xVals = append(xVals, val[0])
    }
    sort.Ints(xVals)
    diff := 0
    for i:=1; i < len(xVals); i++ {
        cur := xVals[i] - xVals[i-1]
        if cur > diff {
            diff = cur
        }
    }
    return diff
}
