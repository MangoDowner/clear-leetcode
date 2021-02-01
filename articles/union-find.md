# 并查集
```go
type UF struct {
    Count  int // 连通分量
	Parent []int // 每个节点的根节点
	Size   []int // 树的重量
}

func NewUF(n int) *UF {
	u := new(UF)
	u.Count = n
	u.Parent = make([]int, n)
	u.Size = make([]int, n)
	for i := 0; i < n; i++ {
		u.Parent[i] = i
		u.Size[i] = 1
	}
	return u
}

// 将p和q连通
func (u *UF) union(p, q int) {
	rootP, rootQ := u.find(p), u.find(q)
	if rootP == rootQ {
		return
	}
	if u.Size[rootP] > u.Size[rootQ] {
		u.Parent[rootQ] = rootP
		u.Size[rootP] += u.Size[rootQ]
	} else {
		u.Parent[rootP] = rootQ
		u.Size[rootQ] += u.Size[rootP]
	}
	u.Count--
}

// 返回节点x的根节点
func (u *UF) find(x int) int {
	for x != u.Parent[x] {
		// 路径压缩
		u.Parent[x] = u.Parent[u.Parent[x]]
		x = u.Parent[x]
	}
	return x
}

// 判断p，q是否连通
func (u *UF) connected(p, q int) bool {
	rootP, rootQ := u.find(p), u.find(q)
    // 处于同一棵输上的节点，相互连通
	return rootP == rootQ
}
```
