# 剑指 Offer 36. 二叉搜索树与双向链表
> 原题链接：[剑指 Offer 36. 二叉搜索树与双向链表](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/)
### 解题思路
由于转换后的双向链表中结点的顺序与二叉树的中序遍历的顺序相同，因此，可以对二叉树的``中序遍历``算法进行修改，
通过在中序遍历的过程中修改结点的指向来转换成一个排序的双向链表。

实现思路如下图所示：
* 1）假设当前遍历的结点为root，root的左子树已经被转换为双向链表（如下图（1）所示）
* 2）使用两个变量``pHead``与``pEnd``分别指向链表的头结点与尾结点
* 3）那么在遍历``root``结点的时候，只需要将``root``结点的``lchild``指向``pEnd``，把``pEnd``的``rchild``（右）指向``root``；
此时root结点就被加入到双向链表里了，因此，root变成了双向链表的尾结点
* 4）对于所有的结点都可以通过同样的方法来修改结点的指向
* 5）因此，可以采用``递归``的方法来求解，在求解的时候需要特别注意递归的结束条件以及边界情况（例如双向链表为空的时候）
### 代码
```go
```
