# 剑指 Offer 40. 最小的k个数
> 原题链接：[剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)
## 1、排序后取前K个数字

## 方法一：排序法
### 解题思路
最简单的方法就是首先对数组进行排序，在排序后的数组中，下标为k-1的值就是第k小的数。
### 复杂度
由于最高效的排序算法（例如快速排序）的平均时间复杂度为O(nlogn)，
因此，此时该方法的平均时间复杂度为O(nlogn)，其中，n为数组的长度。
### 代码
直接用系统的``sort``方法
```go
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
```
## 方法二：部分排序法
### 解题思路
由于只需要找出第k小的数，因此，没必要对数组中所有的元素进行排序，可以采用部分排序的方法

用部分排序的方法。具体思路为：
* 1、通过对选择排序进行改造，第一次遍历从数组中找出最小的数
* 2、第二次遍历从剩下的数中找出最小的数（在整个数组中是第二小的数）
* 3、第k次遍历就可以从n-k+1（n为数组的长度）个数中找出最小的数（在整个数组中是第k小的）

### 复杂度
这种方法的时间复杂度为O(n*k)。当然也可以采用堆排序进行k趟排序找出第k小的值
### 代码
### 计数排序
```go
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


```go
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
## 方法二：部分排序法
### 解题思路
由于只需要找出第k小的数，因此，没必要对数组中所有的元素进行排序，可以采用部分排序的方法

用部分排序的方法。具体思路为：
* 1、通过对选择排序进行改造，第一次遍历从数组中找出最小的数
* 2、第二次遍历从剩下的数中找出最小的数（在整个数组中是第二小的数）
* 3、第k次遍历就可以从n-k+1（n为数组的长度）个数中找出最小的数（在整个数组中是第k小的）

### 复杂度
这种方法的时间复杂度为O(n*k)。当然也可以采用堆排序进行k趟排序找出第k小的值

## 方法三：类快速排序方法
快速排序的基本思想为：
* 1、将数组``array[low…high]``中某一个元素（取第一个元素）作为划分依据，然后把数组划分为三部分
    * 1）``array[low…i-1]``（所有的元素的值都小于或等于``array[i]``）
    * 2）``array[i]``
    * 3）``array[i+1…high]``（所有的元素的值都大于array``[i]``）
* 2、在此基础上可以用下面的方法求出第k小的元素：
    * 1）如果``i-low==k-1``，说明``array[i]``就是第k小的元素，那么直接返回``array[i]``。
    * 2）如果``i-low>k-1``，说明第k小的元素肯定在``array[low…i-1]``中，
    那么只需要递归地在``array[low…i-1]``中找第k小的元素即可。
    * 3）如果``i-low<k-1``，说明第k小的元素肯定在``array[i+1…high]``中，
    那么只需要递归地在``array[i+1…high]``中找第``k-(i-low)-1``小的元素即可。
    
### 代码
```go
//快速选择
func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	res := make([]int, 0)
	left, right := 0, len(arr)-1
	for {
		//基准位置
		pos := partition(arr, left, right)
		//如果基准位置在k处，说明前k个元素就是我们需要的结果
		if pos == k {
			break
		} else if pos < k { //如果小于k，说明所求位置应该在右边
			left = pos + 1
		} else {
			right = pos - 1 //所求位置在左边
		}
	}
	//放入结果集中
	for i := 0; i < k; i++ {
		res = append(res, arr[i])
	}
	return res
}

// 找到pivot位置
func partition(a []int, left, right int) int {
	pivot := left
	for left < right {
		for left < right && a[right] >= a[pivot] {
			right--
		}
		for left < right && a[left] <= a[pivot] {
			left++
		}
		a[right], a[left] = a[left], a[right]
	}
	a[right], a[pivot] = a[pivot], a[right]
	return right
}   
```
