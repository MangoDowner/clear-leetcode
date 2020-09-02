# 剑指 Offer 27. 二叉树的镜像
> 原题链接：[剑指 Offer 27. 二叉树的镜像](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof)
## 递归法
### 解题思路
要实现二叉树的镜像反转，只需交换二叉树中所有结点的左右孩子即可。由于对所有的结点都做了同样的操作，因此，可以用递归的方法来实现，实现代码如下
### 代码
```golang
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
```

## 广度优先遍历
### 解题思路
### 代码
```golang
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		head.Left, head.Right = head.Right, head.Left
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
	return root
}
```