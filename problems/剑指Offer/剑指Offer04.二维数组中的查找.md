# 剑指 Offer 04. 二维数组中的查找
> 原题链接：[剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)
### 解题思路
* 1、从左到右递增的顺序排序、上到下递增的顺序排序
* 2、所以咱们可以从右上角，或者左下角开始遍历
* 3、这次咱们使用从右上角开始遍历，即row从小到大，col从大到小

### 代码
```golang
func findNumberIn2DArray(matrix [][]int, target int) bool {
	rowNum := len(matrix)
	if rowNum == 0 {
		return false
	}
	colNum := len(matrix[0])
	if colNum == 0 {
		return false
	}
	row, col := 0, colNum -1
	for row < rowNum && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
```