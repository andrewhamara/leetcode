func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }

    stack := []*TreeNode{root}

    leftLeavesSum := 0

    for len(stack) > 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if curr.Right != nil {
            stack = append(stack, curr.Right)
        }
        if curr.Left != nil {
            if curr.Left.Left == nil && curr.Left.Right == nil {
                leftLeavesSum += curr.Left.Val
            }else {
                stack = append(stack, curr.Left)
            }
        }
    }
    return leftLeavesSum
}
