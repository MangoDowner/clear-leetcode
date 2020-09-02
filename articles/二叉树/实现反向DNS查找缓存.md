# 实现反向DNS查找缓存

## 题目描述
反向DNS查找指的是使用Internet IP地址查找域名。

例如，如果你在浏览器中输入74.125.200.106，它会自动重定向到对应网址。

如何实现反向DNS查找缓存？

## 分析与解答
要想实现反向DNS查找缓存，主要需要完成如下功能：
* 1）将IP地址添加到缓存中的URL映射。
* 2）根据给定IP地址查找对应的URL

对于本题的这种问题，常见的一种解决方案是使用哈希法（使用map来存储IP地址与URL之间的映射关系），
由于这种方法相对比较简单，这里就不做详细介绍了。

下面重点介绍另外一种方法：
### Trie树
这种方法的主要优点如下：
* 1）使用Trie树，在最坏的情况下的时间复杂度为O(1) ，而哈希方法在平均情况下的时间复杂度为O(1)。
* 2）Trie树可以实现前缀搜索（对于有相同前缀的IP地址，可以寻找所有的URL）

当然，由于树这种数据结构本身的特性，所以使用树结构的一个最大的缺点就是需要耗费更多的内存，
但是对于本题而言，这却不是一个问题，因为Internet IP地址只包含有11个字母（0到9和.）。

所以，本题实现的主要思路为：在Trie树中存储IP地址，而在最后一个结点中存储对应的域名

### 代码
```golang
const CharCount = 11

//Trie树
type TrieNode struct {
	IsLeaf bool
	Url    string
	Child  []*TrieNode
}

func NewTrieNode(count int) *TrieNode {
	return &TrieNode{IsLeaf: false, Url: "", Child: make([]*TrieNode, count)}
}

type DNSCache struct {
	root *TrieNode
}

func NewDNSCache() *DNSCache {
	return &DNSCache{root: NewTrieNode(CharCount)}
}

func (d *DNSCache) getIndexFromRune(char rune) int {
	if char == '.' {
		return 10
	} else {
		return int(char) - '0'
	}
}

func (d *DNSCache) getRuneFromIndex(i int) rune {
	if i == 10 {
		return '.'
	} else {
		return rune('0' + i)
	}
}

// 把一个IP地址和相应的URL添加到trie树中，最后一个节点是URL
func (d *DNSCache) Insert(ip, url string) {
	pCrawl := d.root
	for _, v := range []rune(ip) {
		// 根据当前遍历到的IP中的字符，找到子节点的索引
		index := d.getIndexFromRune(v)
		// 如果子节点不存在，创建一个
		if pCrawl.Child[index] == nil {
			pCrawl.Child[index] = NewTrieNode(CharCount)
		}
		// 移动到子节点
		pCrawl = pCrawl.Child[index]
	}
	// 在叶子节点中存储IP地址对应的URL
	pCrawl.IsLeaf = true
	pCrawl.Url = url
}

// 通过IP地址找到对应的URL
func (d *DNSCache) SearchDNSCache(ip string) string {
	pCrawl := d.root
	for _, v := range []rune(ip) {
		index := d.getIndexFromRune(v)
		if pCrawl.Child[index] == nil {
			return ""
		}
		pCrawl = pCrawl.Child[index]
	}
	// 返回找到的URL
	if pCrawl != nil && pCrawl.IsLeaf {
		return pCrawl.Url
	}
	return ""
}
```
使用范例
```golang
func main() {
	ipAddrs := []string{"10.57.11.127"}
	urls := []string{"www.samsung.com"}
	dnsCache := NewDNSCache()
	for i, v := range ipAddrs {
		dnsCache.Insert(v, urls[i])
	}
	ip := ipAddrs[0]
	result := dnsCache.SearchDNSCache(ip)
	fmt.Println(result)
}
```
