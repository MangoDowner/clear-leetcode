# 剑指 Offer 28. 对称的二叉树

> 原题链接：[剑指 Offer 28. 对称的二叉树](https://leetcode-cn.com/problems/add-two-numbers/)

### 代码
```c#
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     public int val;
 *     public TreeNode left;
 *     public TreeNode right;
 *     public TreeNode(int x) { val = x; }
 * }
 */
public class Solution {
    public bool IsSymmetric(TreeNode root) {
        return isValid(root, root);
    }

    public bool isValid(TreeNode a, TreeNode b)
    {
        if (a == null && b == null)
        {
            return true;
        }
        if (a == null || b == null)
        {
            return false;
        }

        if (a.val != b.val)
        {
            return false;
        }
        return isValid(a.left, b.right) && isValid(a.right, b.left);
    }
}
```