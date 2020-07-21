# 大数越界问题
## 确定最大值后，开辟相应空间数组
### 解题思路
### 代码
```golang
func printNumbers(n int) []int {
	max := 9
	for i := 1; i <n;i++ {
		max = max *10 + 9
	}
	result := make([]int, max)
	for i := 1; i <= max;i ++ {
		result[i-1] = i
	}
	return result
}
```
## 考虑越界问题
### 解题思路
本题主要还是要考察的大数越界问题，所以最好是转为字符串输出。

其中要注意过滤到前置0，比如``008``这样的数字，最好转化为``8``。
### 代码
```golang
var result []string
func printNumbers(n int) []string {
	result = []string{}
	num := make([]string, n)
	dfs(n, 0, -1, num) // start用于过滤前置0
	return result[1:]
}

func dfs(n, x, start int, num []string) {
	if n == x {
		if start != -1 {
			result = append(result, strings.Join(num[start:], ""))
		} else {
			result = append(result, strings.Join(num, ""))
		}
		 return
	}
	for i := 0; i <=9; i++ {
		if start == -1 && i != 0 {
			start = x
		}
		num[x] = strconv.Itoa(i)
		dfs(n, x+1, start, num)
	}
}
```