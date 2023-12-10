func garbageCollection(garbage []string, travel []int) int {
    minutes := 0

    var lastG, lastP, lastM, length int = -1,-1,-1,0

    for i,s := range garbage {
        length += len(s)
        if strings.Contains(s, "G") {
            lastG = i
        }
        if strings.Contains(s, "P") {
            lastP = i
        }
        if strings.Contains(s, "M") {
            lastM = i
        }
    }

    if lastG != -1 {
        for i := 0; i < lastG; i++ {
            minutes += travel[i]
        }
    }
    if lastM != -1 {
        for i := 0; i < lastM; i++ {
            minutes += travel[i]
        }
    }
    if lastP != -1 {
        for i := 0; i < lastP; i++ {
            minutes += travel[i]
        }
    }

    return minutes + length
}
