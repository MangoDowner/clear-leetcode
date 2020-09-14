# 剑指Offer33.二叉搜索树的后序遍历序列

> 原题链接：[剑指 Offer 33. 二叉搜索树的后序遍历序列树](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof)
### 解题思路
二元查找树的特点是：
对于任意一个结点，它的左子树上所有结点的值都小于这个结点的值，它的右子树上所有结点的值都大于这个结点的值。

根据它的这个特点以及二元查找树后序遍历的特点，可以看出，这个序列的最后一个元素一定是树的根结点（上图中的结点4），
然后在数组中找到第一个大于根结点4的值5，那么结点5之前的序列``（1，3，2）``对应的结点一定位于结点4的左子树上，
结点5（包含这个结点）后面的序列一定位于结点``4``的右子树上（也就是说结点5后面的所有值都应该大于或等于4）。

对于结点4的左子树遍历的序列```{1，3，2}``以及右子树的遍历序列``{5，7，6}``可以采用同样的方法来分析。

因此，可以通过递归方法来实现
### 代码
```golang
func verifyPostorder(postorder []int) bool {
	return isPostOrder(postorder, 0, len(postorder)-1)
}

func isPostOrder(arr []int, left, right int) bool {
	if left >= right {
		return true
	}
	// 数组的最后一个节点必定是根节点
	root := arr[right]
	var p, mid int
	// 找到第一个大于root的值，那么前面所有的节点都在位于root的左子树上
	for p = left; arr[p] < root; p++ {
	}
	// 如果序列是后续遍历的序列，那么从i开始的所有值都应该大于根节点root的值
	mid = p
	for ; arr[p] > root; p++ {
	}
	return p == right && isPostOrder(arr, left, mid-1) && isPostOrder(arr, mid, right-1)
}
```