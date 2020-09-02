# 剑指 Offer 24. 反转链表
> 原题链接：[剑指 Offer 24. 反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof)

### 解题思路

### 代码
```golang
func reverseList(head *ListNode) *ListNode {
	var cur, pre, next *ListNode
	cur, next = head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
```