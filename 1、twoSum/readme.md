
##[1] 两数之和
 

###题目
 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返他们的数组下标。
 
 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
 
  示例:
 
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 
###解题思路
1、建立一个map[int]int，用来保存数组的值和索引 map[val]idx
```
    把  a + b = target
    看成 a = target - b 
        a + map[target - b]
```

2、找到数组中间的值，可以从两边向中间查看，只需要循环半次

### LeetCode 结果
✔ Accepted
  ✔ 29/29 cases passed (8 ms)
  ✔ Your runtime beats 95.19 % of golang submissions
  ✔ Your memory usage beats 1.88 % of golang submissions (3.9 MB)