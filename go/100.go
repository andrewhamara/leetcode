func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && isSameTree(q.Left, p.Left) && isSameTree(q.Right, p.Right)
}
