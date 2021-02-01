# 剑指 Offer 56 - I. 数组中数字出现的次数
> 原题链接：[剑指 Offer 56 - I. 数组中数字出现的次数](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)

## 异或法
### 解题思路
* 1、任何一个数字异或它自己其结果都等于0
* 2、所有数字异或后，出现偶数次的数字都抵消了，只剩下两个出现一次的数
* 3、假设两个数字为``a``和``b``，最后异或的结果为``c``
* 4、我们这时候，只要知道c对应的二进制数中某一个位为1的位数``N``，比如十进制``7``二进制表示形式为``111``，那么可取``N=1/2/3``，然后将c与数组中第``N``位为1的数挨个进行异或，异或结果就是a，b中一个，然后用c异或其中一个数，就可以求出另外一个数了。

#### 通过上述方法为什么就能得到问题的解呢？
* 1）``c``中第``N``位为1表示``a``或``b``中有一个数的第``N``位也为1，假设该数为``a``
* 2）如下图所示，当将``c``与数组中第``N``位为1的数进行异或时，也就是将``x``与``a``外加上其他第``N``位为1的出现过偶数次的数进行异或，化简即为``x``与``a``异或，结果即为b

```
c = a ^ b
d ^ d ^ e ^ e ... a ^ c = a ^ c = a ^ a ^ b = b
```
### 代码
```go
func singleNumber(arr []int) []int {
	var result []int
	if arr == nil {
		return result
	}
	xor := 0
	// 计算数组中所有数字异或的结果
	for _, v := range arr {
		xor ^= v
	}
	position := 0
	for i := xor; i&1 == 0; i >>= 1 {
		position++
	}
	// 异或的结果与所有第position位为1的数字异或
	// 结果一定是出现一次的两个数中的一个
	tempXor := xor
	for _, v := range arr {
		if (uint(v)>>uint(position))&1 == 1 {
			xor ^= v
		}
	}
	// 得到另外一个出现一次的数字
	return []int{xor, xor ^ tempXor}
}
```
