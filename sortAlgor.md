## 一、选择排序
### 思路
依次选择数组元素，每次都把后面最小的数换到该位置来

> 选择排序的运行时间和输入的数组是否有序无关

### 代码
```golang
func selectSort(a []int) {
	N := len(a)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			fmt.Println(j)
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}
```

## 二、插入排序
### 思路
从第二个元素开始，依次把数字送到比他大的数字前面

> 当倒置的数量很少时，插入排序可能是最快的排序方法

> 插入排序一般来说比选择排序快很多
### 代码
```golang
func insertSort(a []int) {
	N := len(a)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
```

## 三、希尔排序
基于插入排序，先分段排序，所分段越来越小，保持局部有序，然后实行插入排序
### 代码
```golang
func shellSort(a []int) {
	N, h := len(a), 1
	for h < N/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}
}
```

## 四、快速排序
份儿治之，将小于nums[pivot]的数字放在其左边，大于其的数放在右边

### 代码
```golang
func sort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(a, left, right)
	quickSort(a, pivot+1, right)
	quickSort(a, left, pivot-1)
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

### 改进的快速排序
对于小数组来说，快速排序不如插入排序，那么当left和right距离太近的时候，我们可以采用插入排序。

这个距离太近怎么定义？可以是``5``到``15``中某一个数，这个值和系统相关，不过这个区间内任意一个数字都能取得不错的增益。
```golang
// 快速排序
func quickSort(a []int, left, right int) {
	if left+5 >= right {
		insertSort(a, left, right)
		return
	}
	pivot := partition(a, left, right)
	quickSort(a, pivot+1, right)
	quickSort(a, left, pivot-1)
}

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

// 插入排序
func insertSort(a []int, left, right int) {
	for i := left + 1; i < right; i++ {
		for j := i; j > left && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
```

## 堆排序
### 解题思路
构造优先队列，底层基于数组，每次都把最大值拿到最后面去
### 代码
```golang
func heapSort(arr []int) {
	l := len(arr)
	// 构造大顶堆
	// 此时我们从最后一个非叶子结点开始（叶结点自然不用调整，第一个非叶子结点 arr.length/2-1
	for i := l/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, l)
	}
	// 循环不变量：区间 [0, i] 堆有序
	for i := l - 1; i > 0; i-- {
		// 把堆顶元素（当前最大）交换到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		// 逐步减少堆有序的部分
		// 重新对堆进行调整
		adjustHeap(arr, 0, i)
	}
}

// 调整大顶堆（仅是调整过程，建立在大顶堆已构建的基础上）
func adjustHeap(arr []int, start, length int) {
	temp := arr[start] // 取出当前元素i
	//从i结点的左子结点开始，也就是2i+1处开始
	for i := start*2 + 1; i < length; i = i*2 + 1 {
		//如果左子结点小于右子结点，k指向右子结点
		if i+1 < length && arr[i] < arr[i+1] {
			i++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if arr[i] > temp {
			arr[start] = arr[i] // 将当前的最大值赋值给start位置
			start = i           // 记录当前最大值索引
		} else {
			break
		}
	}
	// 当最开始start位置的值换到最后的子节点去
	arr[start] = temp
}
```
