> 原题链接：[剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
## 1、排序后取前K个数字
### 解题思路
1、调用系统方法sort，从小到大排序
2、返回``arr``前K位
### 代码
直接用系统的``sort``方法
```golang
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
```
或者用``计数排序``
```golang
func getLeastNumbers(arr []int, k int) []int {
	var result []int
	if k == 0 || len(arr) == 0 {
		return result
	}
	counter := make([]int, 10001)
	for _, v := range arr {
		counter[v]++
	}
	index, result := 0, make([]int, k)
	for key := range counter {
		for ; counter[key] > 0 && index < k; index, counter[key] = index+1, counter[key]-1 {
			result[index] = key
		}
		if index == k {
			break
		}
	}
	return result
}
```


```golang
func getLeastNumbers(arr []int, k int) []int {
	l := len(arr)
	var result []int
	if l == 0 {
		return result
	}
	result = append(result, arr[0])
	for i := 1 ; i < l ;i++ {
		t := []int{}
		for len(result) > 0 && result[len(result) -1] > arr[i] {
			t = append([]int{result[len(result) -1]}, t...)
			result = result[:len(result)-1]
		}
		result = append(result, arr[i])
		result = append(result, t...)
	}
	return result[:k]
}
```