# LCP 22. 黑白方格画
> 原题链接：[LCP 22. 黑白方格画](https://leetcode-cn.com/problems/ccw6C7/)

## 题干
小扣注意到秋日市集上有一个创作黑白方格画的摊位。
摊主给每个顾客提供一个固定在墙上的白色画板，画板不能转动。
画板上有 n * n 的网格。绘画规则为，小扣可以选择任意多行以及任意多列的格子涂成黑色，所选行数、列数均可为 0。

小扣希望最终的成品上需要有 k 个黑色格子，请返回小扣共有多少种涂色方案。

注意：两个方案中任意一个相同位置的格子颜色不同，就视为不同的方案。

示例 1：
```
输入：n = 2, k = 2

输出：4

解释：一共有四种不同的方案：
第一种方案：涂第一列；
第二种方案：涂第二列；
第三种方案：涂第一行；
第四种方案：涂第二行。
```
示例 2：
```
输入：n = 2, k = 1

输出：0

解释：不可行，因为第一次涂色至少会涂两个黑格。
```
示例 3：
```
输入：n = 2, k = 4

输出：1

解释：共有 2*2=4 个格子，仅有一种涂色方案。
```

限制：
```
1 <= n <= 6
0 <= k <= n * n
```

### 解题思路
+ k == 0 表示不涂色，此时方案就一种
+ k < n 无法满足，任意涂一行或者一列产生的格子数最小都是n
+ k == n * n表示全部涂色，此时方案也是一种
其他情况枚举涂色的行数与列数，比如当前涂了i行和j列，此时黑色的方格数量为 
``n * i + n * j - i * j``如果此时的数量恰好等于k，说明当前这种组合方案满足题目，将其累加到答案中。

组合数公式利用递推公式来求解：``C[m][n] = C[m - 1][n - 1] + C[m - 1][n]``


### 代码
```go
func paintingPlan(n int, k int) int {
	c := make([][]int, 7)
	for i := 0; i <= 6; i++ {
		c[i] = make([]int, 7)
		for j := 0; j <= i; j++ {
			if j == 0 {
				c[i][j] = 1
			} else {
				c[i][j] = c[i-1][j-1] + c[i-1][j]
			}
		}
	}

	if k == 0 || k == n*n {
		return 1
	}
	if k < n {
		return 0
	}
	result := 0
	for rowNum := 0; rowNum <= n; rowNum++ {
		for colNum := 0; colNum <= n; colNum++ {
			if n*rowNum+n*colNum-rowNum*colNum == k {
				result += c[n][rowNum] * c[n][colNum]
			}
		}
	}
	return result
}
```
