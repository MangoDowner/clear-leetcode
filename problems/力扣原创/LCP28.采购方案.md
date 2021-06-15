# LCP 28. 采购方案
小力将 N 个零件的报价存于数组 nums。小力预算为 target，假定小力仅购买两个零件，
要求购买零件的花费不超过预算，请问他有多少种采购方案。

注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1

示例 1：
```
输入：nums = [2,5,3,5], target = 6

输出：1

解释：预算内仅能购买 nums[0] 与 nums[2]。
```
示例 2：
```
输入：nums = [2,2,1,9], target = 10

输出：4

解释：符合预算的采购方案如下：
nums[0] + nums[1] = 4
nums[0] + nums[2] = 3
nums[1] + nums[2] = 3
nums[2] + nums[3] = 10
```
提示：
```
2 <= nums.length <= 10^5
1 <= nums[i], target <= 10^5
```
## 双指针
### 解题思路
+ 1、先给数组从小到大排序
+ 2、左右指针分别指向数组头尾
+ 3、如果当前头尾数字和<=目标值target，那么left和right中间的数字，都可以和left配成一对，所以计数增加``right-left``，
然后右移左指针，即``left++``
+ 4、如果当前头尾数字和>目标值target，那么``和``需要缩小，就把右指针往左移动，即``right--``
### 代码

```csharp
public class Solution {
    public int PurchasePlans(int[] nums, int target)
    {
        Array.Sort(nums);
        const int mod = 1000000007;
        var result = 0;
        int left = 0, right = nums.Length - 1;
        while (left < right)
        {
            if (nums[left] + nums[right] <= target)
            {
                result = (result + right - left) % mod;
                left++;
            }
            else
            {
                --right;
            }
        }

        return result;
    }
}
```
