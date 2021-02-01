# 剑指 Offer 66. 构建乘积数组
> 原题链接：[剑指 Offer 66. 构建乘积数组](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof)

## 前后缀乘积
### 解题思路
* 1、从前往后遍历，求出每个数的前缀乘积
* 2、从后往前遍历，求出每个数的前缀乘积
* 3、再次遍历数组，从每个数的前缀积乘以后缀积，就是我们要的结果了
### 代码
```go
func constructArr(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	left, right := make([]int, l), make([]int, l)
	left[0], right[l-1] = 1, 1
	// 计算前缀乘积
	for i := 1; i < l; i++ {
		left[i] = left[i-1] * a[i-1]
	}
	// 计算前缀乘积
	for i := l - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}
```
或者改进一下，减少一次循环
```go
func constructArr(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}
	result := make([]int, l)
	result[0] = 1
	// 计算前缀乘积
	for i := 1; i < l; i++ {
		result[i] = result[i-1] * a[i-1]
	}
	// 直接将后缀乘积乘到最终结果上
	result[0] = a[l-1] // 用于暂存右边数字的后缀乘积
	for i := l - 2; i >= 1; i-- {
		result[i] *= result[0]
		result[0] *= a[i]
	}
	return result
}
```
