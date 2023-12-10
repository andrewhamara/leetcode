import (
    "strconv"
)

func calPoints(operations []string) int {

    scores := make([]int, 0)

    for _,op := range operations {
        switch op{
        case "+":
            scores = append(scores, scores[len(scores) - 1] + scores[len(scores) - 2])
        case "D":
            scores = append(scores, scores[len(scores) - 1] * 2)
        case "C":
            scores = scores[:len(scores) - 1]
        default:
            num,_ := strconv.Atoi(op)
            scores = append(scores, num)
        }
    }

    sum := 0
    for _,n := range scores {
        sum += n
    }
    return sum
}
