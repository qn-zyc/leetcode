# divide_two_integers

Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.



来源：<https://discuss.leetcode.com/topic/15568/detailed-explained-8ms-c-solution>

假设 `15/3`，15是被除数，3是除数，除法就是让我们找出需要用多少次从被除数中减去除数但被除数不能为负数。

让我们开始吧，`15-3=12`, 12是正的。让我们减的更多。我们把3左移一位得到6（`3<<1 = 6`），15-6依然得到正的结果。我们再左移一次得到12（`3<<2=12`）,15-12依然得到正的结果。我们再次左移并且得到24（`3<<3 = 24`），我们知道我们最多只能减去12. 把3左移两次得到12，`12 = 3 * 4`，那么我们怎么得到4呢？从1开始左移两次就可以得到4.我们把4累加到答案上（初始为0）。上面的处理过程类似于`15 = 3 * 4 + 3`，现在我们得到了商4和余数3.

我们再次重复这个过程，我们从余数中减去被除数并且得到0（`3 - 3 = 0`），我们知道我们完成了。没有左移，所以我们简单地将 `1 << 0` 累加到答案上。

现在我们有了完整的算法实现除法。

我们需要处理一些异常情况，比如下面列出的：

1. 除数为零
2. 被除数是 `INT_MIN`, 除数是 `-1`（因为 `abs(INT_MIN) = INT_MAX + 1`）

当然我们还需要处理符号，这个就很简单了。