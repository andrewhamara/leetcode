func reverseList(head *ListNode) *ListNode {

    var prev *ListNode = nil
    var curr *ListNode  = head

    for curr != nil {
        var next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
