# 剑指 Offer 53 - I. 在排序数组中查找数字 I
>> 二分法找到左右边界

> 原题链接：[剑指 Offer 53 - I. 在排序数组中查找数字 I](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)

### 解题思路
1、利用``二分法``分别找出``target``数字出现在数组的左右边界
2、``右边界-左边界+1``就是咱们要的最终结果
### 代码
```go
func search(nums []int, target int) int {
	return rightBound(nums, target) - leftBound(nums, target) + 1
}

// 寻找数字的左边界
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid -1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	return left
}

// 寻找数字的右边界
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid -1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	return right
}
```
