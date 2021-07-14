# 剑指 Offer 55 - I. 二叉树的深度
### 代码
c#
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
    public int MaxDepth(TreeNode root) {
        if (root == null)
        {
            return 0;
        }

        var left = MaxDepth(root.left);
        var right = MaxDepth(root.right);
        return Math.Max(left, right) + 1;
    }
}
```