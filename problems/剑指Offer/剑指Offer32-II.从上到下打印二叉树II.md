# 剑指 Offer 32 - II. 从上到下打印二叉树 II
> 原题链接：[剑指 Offer 32 - II. 从上到下打印二叉树 II](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/)
## 广度优先遍历
### 解题思路
如题，用一个队列，一层一层遍历就好了
### 代码
```go
 levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		var curLevel []int
		for i := 0; i < size; i++ {
			node := queue[i]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		result = append(result, curLevel)
	}
	return result
}
```
