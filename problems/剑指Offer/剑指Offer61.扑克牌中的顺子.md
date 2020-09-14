# 剑指 Offer 61. 扑克牌中的顺子
> 原题链接：[剑指 Offer 61. 扑克牌中的顺子](https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/)

## 统计万能牌和空缺
### 解题思路
* 1、给牌从小到大排序
* 2、统计``joker``牌，用来弥补空缺的数字
* 3、统计空缺的数数量，也就是相隔超过1的两张牌中间的牌，比如``2``和``4``之间，缺了``3``
* 4、比较``joker``牌数量是否大于等于空缺牌数量，是的话，就可以用万能的``joker``牌来顶替缺失牌
### 代码
```golang
public class Solution
{
    public bool IsStraight(int[] nums)
    {
        if (nums == null || nums.Length == 0)
        {
            return false;
        }
        int missing = 0, joker = 0;
        Array.Sort(nums);
        while (nums[joker] == 0)
        {
            joker++;
        }
        for (var index = joker + 1; index < nums.Length; index++)
        {
            if (nums[index] == nums[index - 1])
            {
                return false;
            }
            missing += nums[index] - nums[index - 1] - 1;
        }
        return joker >= missing;
    }
}
```
