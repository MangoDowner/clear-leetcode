# 剑指 Offer II 032. 有效的变位词

给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。

注意：若s 和 t中每个字符出现的次数都相同且字符顺序不完全相同，则称s 和 t互为变位词（字母异位词）。

示例1:
```
输入: s = "anagram", t = "nagaram"
输出: true
```
示例 2:
```
输入: s = "rat", t = "car"
输出: false
```
示例 3:
```
输入: s = "a", t = "a"
输出: false
```

提示:
```
1 <= s.length, t.length <= 5 * 104
sandt仅包含小写字母
```

进阶:如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

注意：本题与主站 242题相似（字母异位词定义不同）：https://leetcode-cn.com/problems/valid-anagram/
```
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dKk3P7
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```
## 哈希表法
### 解题思路
+ 用一个哈希表，计算两个字符串中的字符字数，看看是否一致即可
+ 也可以简化下，一个加，一个减，看看是否能中和就完事了
### 代码

```golang
func isAnagram(s string, t string) bool {
	// 完全相同或者长度不一，不是
	if s == t || len(s) != len(t) {
		return false
	}
	m := make(map[uint8]int)
	for i := 0; i < len(s);i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
```