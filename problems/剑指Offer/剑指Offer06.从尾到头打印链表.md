# 剑指 Offer 06. 从尾到头打印链表
> 原题链接：[剑指 Offer 06. 从尾到头打印链表](https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/)

### 代码
c#
```csharp
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int x) { val = x; }
 * }
 */
public class Solution {
    public int[] ReversePrint(ListNode head)
    {
        var ret = Recur(head);
        return ret.ToArray();
    }

    private List<int> Recur(ListNode head)
    {
        if (head == null)
        {
            return new List<int>();
        }

        var ret = Recur(head.next);
        ret.Add(head.val);
        return ret;
    }
}
```