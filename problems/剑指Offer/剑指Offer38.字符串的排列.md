# 剑指 Offer 38. 字符串的排列
> 原题链接：[剑指 Offer 38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)

## 回溯法
### 解题思路
以下以字符串``abc``为例介绍对字符串进行全排列的方法。具体步骤如下：
* 1、首先固定第一个字符``a``，然后对后面的两个字符b与c进行全排列。
* 2、交换第一个字符与其后面的字符，即交换``a``与``b``，然后固定第一个字符``b``，接着对后面的两个字符a与c进行全排列
* 3、由于第2步交换了``a``和``b``破坏了字符串原来的顺序，因此，需要再次交换``a``和``b``使其恢复到原来的顺序，
然后交换第一个字符与第三个字符（交换``a``和``c``），接着固定第一个字符``c``，对后面的两个字符``a``与``b``进行全排列。
在对字符串求全排列的时候就可以采用递归的方式来求解

### 剪枝
当字符串存在重复字符时，排列方案中也存在重复方案。

为排除重复方案，需在固定某位字符时，保证 “每种字符只在此位固定一次” ，
即遇到重复字符时不交换，直接跳过。从 DFS 角度看，此操作称为 “剪枝” 

### 代码
```go
var result []string

func permutation(s string) []string {
	result = []string{}
	str := []rune(s)
	dfs(str,0)
	return result
}

func dfs(str []rune, start int) {
	if start == len(str) - 1 {
		result = append(result, string(str))
		return
	}
	exist := make(map[rune]bool)
	for i := start; i < len(str); i++ {
		if exist[str[i]] {
			continue
		}
		exist[str[i]] = true
		str[start], str[i] = str[i], str[start]
		dfs(str, start+1)
		str[start], str[i] = str[i], str[start]
	}
}
```
