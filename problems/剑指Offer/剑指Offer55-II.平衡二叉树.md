# 剑指 Offer 55 - II. 平衡二叉树

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
public class Solution
{
    public bool IsBalanced(TreeNode root)
    {
        if (root == null)
        {
            return true;
        }

        if (Math.Abs(depth(root.left) - depth(root.right)) >= 2)
        {
            return false;
        }

        return IsBalanced(root.left) && IsBalanced(root.right);
    }

    public int depth(TreeNode root)
    {
        if (root == null)
        {
            return 0;
        }

        var left = depth(root.left);
        var right = depth(root.right);
        return Math.Max(left, right) + 1;
    }
}
```