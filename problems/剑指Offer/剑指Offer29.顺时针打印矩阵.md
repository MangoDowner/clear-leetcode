# 浅显易懂且双百
### 解题思路
就按照人划线的思路来做就好了：

![1](../../pictures/problems/offer29/1.png)

1、运动方向分成四个：
```
1、向右
2、向下
3、向左
4、向上
```
向右运动结束，就会自动向下运动，接着向左，向上，这个很好理解

2、每次运动完后，边界都会发生变化
```
1、向右运动结束后，上边界下移，即top++
2、向下运动结束后，右边界左移，即right--
3、向左运动结束后，下边界上移，即bottom--
4、向上运动结束后，左边界右移，即left++
```

3、 每次运动结束后，开始新方向运动时，对应的开始行列也会发生变化
```
1、向右运动结束后，开始行下移，即row++
2、向下运动结束后，开始列左移，即col--
3、向左运动结束后，开始行上移，即row--
4、向上运动结束后，开始列右移，即col++
```

运动结束的时候，也就是上下边界 / 左右边界 重叠的时候。
### 代码
按照这个思路写出代码就好啦！
```go
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
		return []int{}
	}
	dir := 1
	row, col := 0, 0
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var res []int
	for top <= bottom && left <= right {
		res = append(res, matrix[row][col])
		switch dir {
		case 1:
			if col == right {
				dir = 2
				top++
				row++
				continue
			}
			col++
		case 2:
			if row == bottom {
				dir = 3
				right--
				col--
				continue
			}
			row++
		case 3:
			if col == left {
				dir = 4
				bottom--
				row--
				continue
			}
			col--
		case 4:
			if row == top {
				dir = 1
				left++
				col++
				continue
			}
			row--
		}
	}
	return res
}
```
