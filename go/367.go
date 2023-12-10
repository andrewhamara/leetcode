import (
    "fmt"
    "math"
)

func isPerfectSquare(num int) bool {
    root := math.Sqrt(float64(num))
    return root == math.Floor(root)
}
