# 剑指 Offer II 069. 山峰数组的顶部

符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
+ arr.length >= 3
+ 存在 i（0 < i< arr.length - 1）使得： 
  + arr[0] < arr[1] < ... arr[i-1] < arr[i]
  + arr[i] > arr[i+1] > ... > arr[arr.length - 1]

给定由整数组成的山峰数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i，
即山峰顶部。



示例 1：
```
输入：arr = [0,1,0]
输出：1
```
示例 2：
```
输入：arr = [1,3,5,4,2]
输出：2
```
示例 3：
```
输入：arr = [0,10,5,2]
输出：1
```
示例 4：
```
输入：arr = [3,4,5,1]
输出：2
```
示例 5：
```
输入：arr = [24,69,100,99,79,78,67,36,26,19]
输出：2
```

提示：
```
3 <= arr.length <= 104
0 <= arr[i] <= 106
题目数据保证 arr 是一个山脉数组
```

进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？

```
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/B1IidL
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```

## 二分法
### 解题思路
+ 1、设定arr左右边界，left为开始元素，right为最后元素
+ 2、取left和right中间的元素mid，如果中间元素大于其两边邻居元素，那么该mid就是我们要的值
+ 3、如果mid元素和其左右邻居正好递增，那么说明山峰在更右边的位置，left = mid
+ 4、与3情况类似，如果mid元素和其左右邻居正好递减，那么说明山峰在更左边的位置，right = mid
### 代码

```csharp
public class Solution {
    public int PeakIndexInMountainArray(int[] arr) {
        int left = 0, right = arr.Length -1;
        while(left < right) {
            var mid = (left + right) / 2;
            if (arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1]) {
                return mid;
            } 
            if (arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1]) {
                left = mid;
            } else if (arr[mid] < arr[mid-1] && arr[mid] > arr[mid+1]) {
                right = mid;
            }
        }

        return 0;
    }
}
```