# 剑指 Offer 10- II. 青蛙跳台阶问题
> 原题链接：[剑指 Offer 10- II. 青蛙跳台阶问题](https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/)

## 迭代法
### 解题思路
* 1、看到这题就想起``爬楼梯``，就想起``斐波那契数列``
* 3、对于最后一步来说，咱们可以:
    * 1）从倒数第二级走1步上来
    * 2）从倒数第三级走2步上来
* 4、所有最后一步的可能走法取决于到前面两级台阶的走法
* 5、确定base case,即只有``0/1/2``级台阶时候的走法，分别为``0/1/2``种走法
### 代码
```golang
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, (one + two) % 1000000007
	}
	return two
}
```
