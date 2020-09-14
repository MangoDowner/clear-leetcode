# 剑指 Offer 53 - II. 0～n-1中缺失的数字
> 原题链接：[剑指 Offer 53 - II. 0～n-1中缺失的数字](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)

### 解题思路
```golang
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		num := nums[mid]
		if num == mid {
			left = left + 1
		} else if num > mid {
			right = mid - 1
		}
	}
	return left
}
```