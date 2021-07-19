# 剑指 Offer 56 - II. 数组中数字出现的次数 II
> 原题链接：[剑指 Offer 56 - II. 数组中数字出现的次数 II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)
### 解题思路
二进制只能表示 ``0, 10,1`` ，因此需要使用两个二进制位来表示 ``3``个状态。
设此两位分别为 ``two`` , ``one`` ，则状态转换变为：
```
00→01→10→00→⋯
```

接下来，需要通过 状态转换表 导出 状态转换的计算公式 。首先回忆一下位运算特点，对于任意二进制位 xx ，有：
```
异或运算：x ^ 0 = x​ ， x ^ 1 = ~x
与运算：x & 0 = 0 ， x & 1 = x
```
#### 计算 one 方法：

设当前状态为 ``two one`` ，此时输入二进制位 ``n`` 。如下图所示，通过对状态表的情况拆分，
可推出 ``one`` 的计算方法为：

```
if two == 0:
  if n == 0:
    one = one
  if n == 1:
    one = ~one
if two == 1:
    one = 0
```
引入 ``异或运算`` ，可将以上拆分简化为：
```
if two == 0:
    one = one ^ n
if two == 1:
    one = 0
```    
引入 ``与运算`` ，可继续简化为：

```
one = one ^ n & ~two
```
#### 计算 two 方法：

由于是先计算 ``one`` ，因此应在新 ``one``的基础上计算 ``two`` 。
如下图所示，修改为新 ``one`` 后，得到了新的状态图。观察发现，可以使用同样的方法计算 ``two`` ，即：

```
two = two ^ n & ~one
```

### 代码
```go
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, v := range nums {
		ones = ones ^ v & ^twos
		twos = twos ^ v & ^ones
	}
	return ones
}
```
