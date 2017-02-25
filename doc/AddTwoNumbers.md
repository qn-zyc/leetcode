# Add Two Numbers

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Input:** (2 -> 4 -> 3) + (5 -> 6 -> 4)
**Output:** 7 -> 0 -> 8



意思是 342 + 465 = 807 (顺序在链表里是倒序)



思路：用一个变量sum记录相加的和，每前进一步，将两个链表对应节点的值与sum相加，将sum的个位数存放到新节点中，sum缩小10倍，继续。