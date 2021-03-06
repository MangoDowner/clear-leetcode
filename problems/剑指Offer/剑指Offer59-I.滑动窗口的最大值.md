# 剑指 Offer 59 - I. 滑动窗口的最大值
> 原题链接：[剑指 Offer 59 - I. 滑动窗口的最大值](https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/)

## 单调队列

单调队列是⼀个「队列，队列中的元素单调递增（或递减）

* 1、维护窗口，向右移动时左侧超出窗口的值弹出
* 2、因为需要的是窗口内的最大值，所以只要保证窗口内的值是递减的即可，小于新加入的值全部弹出。
* 3、最左端即为窗口最大值
```go
type MonotonicQueue struct {
	Deque []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{}
}

func (m *MonotonicQueue) Push(num int) {
	for len(m.Deque) != 0 && m.Deque[len(m.Deque)-1] < num {
		m.Deque = m.Deque[:len(m.Deque)-1]
	}
	m.Deque = append(m.Deque, num)
}

func (m *MonotonicQueue) Pop(num int) {
	if len(m.Deque) != 0 && m.Deque[0] == num {
		m.Deque = m.Deque[1:]
	}
}

func (m *MonotonicQueue) Max() int {
	return m.Deque[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	var window MonotonicQueue
	var result []int
	for key, value := range nums {
		if key < k-1 { //先填满窗⼝的前 k - 1
			window.Push(value)
		} else {  // 窗⼝向前滑动
			window.Push(value)
			result = append(result, window.Max())
			window.Pop(nums[key-k+1])
		}
	}
	return result
}
```
