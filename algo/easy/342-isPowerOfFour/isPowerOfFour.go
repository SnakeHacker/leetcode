package isPowerOfFour

/*
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

输入: 16
输出: true
示例 2:

输入: 5
输出: false

进阶：
你能不使用循环或者递归来完成本题吗？
*/

/*
@Author: Mickey
*/

func isPowerOfFour(num int) bool {
	// num 非负数 && 只有首位为1 && 偶数位为0
	return num > 0 && (num&(num-1) == 0) && (num&0xaaaaaaaa == 0)
}
