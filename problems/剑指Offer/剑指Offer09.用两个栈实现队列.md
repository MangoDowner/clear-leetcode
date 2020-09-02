# 剑指 Offer 09. 用两个栈实现队列

> 原题链接：[剑指 Offer 09. 用两个栈实现队列](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

本题类似 [232. 用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/)

## 解题思路
* 1、使用栈A与栈B模拟队列Q，A为插入栈，B为弹出栈，以实现队列Q。
* 2、假设A和B都为空，可以认为栈A提供入队列的功能，栈B提供出队列的功能。
* 3、要入队列，入栈A即可，而出队列则需要分两种情况考虑：
    * （1）如果栈B不为空，则直接弹出栈B的数据。
    * （2）如果栈B为空，则依次弹出栈A的数据，放入栈B中，再弹出栈B的数据。
### 代码
```golang
type CQueue struct {
	write []int
	read []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int)  {
	this.write = append(this.write, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.read) == 0 {
		for len(this.write) != 0 {
			this.read = append(this.read, this.write[len(this.write)-1])
			this.write = this.write[:len(this.write)-1]
		}
	}
	if len(this.read) == 0 {
		return -1
	}
	result := this.read[len(this.read)-1]
	this.read = this.read[:len(this.read)-1]
	return result
}
```