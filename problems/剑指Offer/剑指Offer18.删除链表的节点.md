# 剑指 Offer 18. 删除链表的节点

> 原题链接：[剑指 Offer 18. 删除链表的节点](https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/)


```golang
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head!= nil{
		if head.Val == val{
			pre.Next = head.Next
			break
		}
		pre, head = head, head.Next
	}
	return dummy.Next
}
```
或者
```golang
func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
```

## 扩展：如何在只给定单链表中某个结点指针的情况下删除该结点
一般而言，要删除单链表中的一个结点p，首先需要找到结点p的前驱结点pre，
然后通过``pre.next=p.next``来实现对结点p的删除。

对于本题而言，由于无法获取到结点p的前驱结点，因此，不能采用这种传统的方法。

那么如何解决这个问题呢？可以分如下两种情况来分析：
* （1）如果这个结点是链表的最后一个结点，那么无法删除这个结点。
* （2）如果这个结点不是链表的最后一个结点，可以通过把其后继结点的数据复制到当前结点中，然后删除后继结点的方法来实现。
### 代码
```golang
func deleteNode(node *ListNode) bool {
	if node == nil || node.Next == nil {
		return false
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return true
}
```