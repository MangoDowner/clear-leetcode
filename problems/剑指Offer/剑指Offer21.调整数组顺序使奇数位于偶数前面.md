# 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
> 原题链接：[剑指 Offer 21. 调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/)

## 双指针法
###
### 代码
```go
func exchange(nums []int) []int {
	left, right := 0, len(nums) - 1
	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		for left < len(nums) && nums[left]%2 != 0 {
			left++
		}
		for right >= 0 && nums[right]%2 == 0 {
			right--
		}
	}
	return nums
}
```
