# Combination Sum 2

来源：<https://leetcode.com/problems/combination-sum-ii/>

Given a collection of candidate numbers (**C**) and a target number (**T**), find all unique combinations in **C** where the candidate numbers sums to **T**.

Each number in **C** may only be used **once** in the combination.

**Note:**

- All numbers (including target) will be positive integers.
- The solution set must not contain duplicate combinations.

For example, given candidate set `[10, 1, 2, 7, 6, 1, 5]` and target `8`, 
A solution set is: 

```
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

给定一个数组C和一个数字T,找出所有的组合使其和等于T（C中的数字和T都是正数，找到的任意两个组合不能重复）。

C中的数字只可以使用一次。

例如：给定 `C = [10, 1, 2, 7, 6, 1, 5]` 和 `T = 8`， 所有组合为：

```
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

