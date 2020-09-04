# 剑指Offer58-II.左旋转字符串.md
> 原题链接：[剑指 Offer 58 - II. 左旋转字符串](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)

### 解题思路
三次翻转
### 代码
C#
```c#
public class Solution
{
    public string ReverseLeftWords(string s, int n)
    {
        s = ReverseFromTo(s, 0, s.Length - 1);
        s = ReverseFromTo(s, 0, s.Length - n - 1);
        s = ReverseFromTo(s, s.Length - n, s.Length - 1);
        return s;
    }
    private static string ReverseFromTo(string s, int left, int right)
    {
        var sArr = s.ToCharArray();
        char temp;
        while (left < right)
        {
            temp = sArr[left];
            sArr[left] = sArr[right];
            sArr[right] = temp;
            left++;
            right--;
        }

        return new string(sArr);
    }
}
```