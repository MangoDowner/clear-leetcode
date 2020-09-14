# 剑指 Offer 57 - II. 和为s的连续正数序列
>> 滑动窗口法

> 原题链接：[剑指 Offer 57 - II. 和为s的连续正数序列](https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/)

## 滑动窗口法
### 解题思路
* 1、设定左右指针``left``/``right``，分别指向``1``和``2``，维护左右指针间的窗口和``sum``
* 2、如果``sum``等于``target``值，将左右指针间的数字作为一个数组加入最终结果，同时右移右指针
* 3、如果``sum``大于``target``值，缩小窗口，通过右移左指针，将左边的数字移出窗口
* 4、如果``sum``小于``target``值，增大窗口，通过右移右指针，将右边的数字加入窗口
* 5、左指针最多能移动到``target+1/2``的地方，因为至少需要两个数字，那么最大也就是中间的数字了
### 代码
```golang
func findContinuousSequence(target int) [][]int {
	var result [][]int
	if target < 3 {
		return result
	}
	left, right, mid := 1, 2, (target + 1)/2
	sum := left + right
	for left < mid {
		if sum == target {
			// 组装可用数组
			var temp []int
			for i := left; i <= right; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		} else {
			right++
			sum += right
		}
	}
	return result
}
```