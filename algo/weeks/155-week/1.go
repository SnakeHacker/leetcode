package week

/*

请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数的索引」上；你需要返回可能的方案总数。

让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。

由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。

示例 1：

输入：n = 5
输出：12
解释：举个例子，[1,2,5,4,3] 是一个有效的排列，但 [5,2,3,4,1] 不是，因为在第二种情况里质数 5 被错误地放在索引为 1 的位置上。
示例 2：

输入：n = 100
输出：682289015
*/

func numPrimeArrangements(n int) int {
	primeNum := 0

	for i := 1; i < n+1; i++ {
		if isPrime(i) {
			primeNum++
		}
	}

	primeChoice := factorial(primeNum)
	restChoice := factorial(n - primeNum)

	return int((primeChoice * restChoice) % uint64(1000000007))
}

func factorial(n int) uint64 {
	if n > 0 {
		return (uint64(n) * factorial(n-1)) % uint64(1000000007)
	}
	return 1
}

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
