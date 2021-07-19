# 剑指 Offer 37. 序列化二叉树
> 原题链接：[剑指 Offer 37. 序列化二叉树](https://leetcode-cn.com/problems/xu-lie-hua-er-cha-shu-lcof)

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
public class Codec {
    // Encodes a tree to a single string.
    public string serialize(TreeNode root)
    {
        if (root == null) return "null,";
        string res = root.val + ",";
        res += serialize(root.left);
        res += serialize(root.right);
        return res;
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(string data)
    {
        string[] arr = data.Split(',');
        Queue<string> queue = new Queue<string>();
        for (int i = 0; i < arr.Length; i++) 
        {
            queue.Enqueue(arr[i]);
        }
        return Helper(queue);
    }

    public TreeNode Helper(Queue<string> queue)
    {
        string val = queue.Dequeue();
        if (val == "null") 
        {
            return null;
        }
        TreeNode root = new TreeNode(Convert.ToInt32(val));
        root.left = Helper(queue);
        root.right = Helper(queue);
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
```
