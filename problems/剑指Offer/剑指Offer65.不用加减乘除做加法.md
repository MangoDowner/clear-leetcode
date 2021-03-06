# 剑指 Offer 65. 不用加减乘除做加法
> 原题链接：[剑指 Offer 65. 不用加减乘除做加法](https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/)

### 解题思路
由于不能使用加减乘除运算，因此，只能使用位运算了。
首先通过分析十进制加法的规律来找出二进制加法的规律，从而把加法操作转换为二进制的操作来完成。

十进制的加法运算过程可以分为以下三个步骤：
* 1、各个位相加而不考虑进位，计算相加的结果sum；
* 2、只计算各个位相加时进位的值carry；
* 3、将sum与carry相加就可以得到这两个数相加的结果。例如15+29的计算方法为：sum=34（不考虑进位），
carry=10（只计算进位），因此，``15+29=sum+carry=34+10=44``。
同理，二进制加法与十进制加法有着相似的原理，
唯一不同的是，在二进制加法中，``sum``与``carry``的和可能还有进位，
因此，在二进制加法中会不停地执行``sum+carry``操作，直到没有进位为止。

具体实现方法如下：
* 1、二进制各个位相加而不考虑进位。由于在不考虑进位的时候加法操作可以用异或操作代替，
因此，不考虑进位的加法可以用异或运算来代替。
* 2、计算进位，由于只有1+1才会产生进位，因此，进位的计算可以用与操作代替。
进位的计算方法为：先做与运算，再把运算结果左移一位。
* 3、不断对``1/2``两步得到的结果相加，直到进位为0的时候为止。

### 代码
```go
func add(a int, b int) int {
	sum := 0 // 保存不进位的相加结果
	carry := -1 // 保存进位值
	for carry != 0 {
		sum = a ^ b // 异或代替不进位相加
		carry = (a&b) << 1 // 与操作代替计算进位值
		a, b = sum, carry
	}
	return sum
}
```

### 引申：如何不用加减乘除运算实现减法
由于减去一个数等于加上这个数的相反数，即
```-n=～(n-1)=～n+1``，因此，``a-b=a+(-b)=a+(～b)+1``，可以利用上面已经实现的加法操作来实现减法操作
### 代码
```go
func sub(a, b int) int {
	return add(a, add(^b, 1))
}

func add(a int, b int) int {
	sum := 0 // 保存不进位的相加结果
	carry := -1 // 保存进位值
	for carry != 0 {
		sum = a ^ b // 异或代替不进位相加
		carry = (a&b) << 1 // 与操作代替计算进位值
		a, b = sum, carry
	}
	return sum
}
```
