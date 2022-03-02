# 剑指 Offer II 024. 反转链表

给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。

示例 1：
```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
```
示例 2：
```
输入：head = [1,2]
输出：[2,1]
```
示例 3：
```
输入：head = []
输出：[]
```

提示：
```
链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
```

进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？


注意：本题与主站 206题相同：https://leetcode-cn.com/problems/reverse-linked-list/
```
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/UHnkqh
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

## 迭代
### 解题思路
见代码
### 代码
```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
```
## 递归
### 解题思路
见代码

### 代码

```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```