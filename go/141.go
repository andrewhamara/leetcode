// Author: Andrew Hamara

// Solution to leetcode problem 141. Linked List Cycle

func hasCycle(head *ListNode) bool {

    hare := head
    tortoise := head

    for hare != nil && tortoise != nil {
        hare = hare.next.next
        tortoise = tortoise.next

        if hare == tortoise {
            return true
        }
    }
    return false
}
