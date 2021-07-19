# 剑指 Offer 57. 和为s的两个数字
> 原题链接：[剑指 Offer 57. 和为s的两个数字](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/)
## 双指针法
### 解题思路
* 1、初始化： 双指针 ``i`` , ``j`` 分别指向数组 ``nums`` 的左右两端 （俗称对撞双指针）。
* 2、循环搜索： 当双指针相遇时跳出；
    * 1）计算和 ``s = nums[i] + nums[j]`` ；
    * 2）若 ``s > targets`` ，则指针 ``j`` 向左移动，即执行 ``j = j - 1`` ；
    * 3）若 ``s < targets`` ，则指针 ``i`` 向右移动，即执行 ``i = i + 1`` ；
    * 4）若 ``s = targets`` ，立即返回数组 ``[nums[i], nums[j]]`` ；
* 3、返回空数组，代表无和为 ``target`` 的数字组合。

### 代码
```go
func twoSum(nums []int, target int) []int {
	var result []int
	l := len(nums)
	if l <= 1 {
		return result
	}
	left, right := 0, l-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return result
}
```
