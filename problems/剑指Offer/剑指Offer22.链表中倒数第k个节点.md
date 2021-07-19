# 剑指 Offer 22. 链表中倒数第k个节点
> ## 快指针先行k步

> [剑指 Offer 22. 链表中倒数第k个节点](https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/)
 
## 快慢指针法
### 解题思路
由于单链表只能从头到尾依次访问链表的各个结点，因此，如果要找链表的倒数第k个元素，
也只能从头到尾进行遍历查找

* 1、在查找过程中，设置快慢两指针，让快指针先行``k``步
* 2、然后两个指针同时往前移动
* 3、循环直到快指针值为``nil``时，慢指针所指的位置就是所要找的位置

其中，要注意，链表可能没有``k``那么长，那么倒数第``k``位，就是``nil``了

### 代码
```go
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	i := 0
	// 快指针先走k步
	for ; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	// 没走完k步
	if i < k {
		return nil
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
```
