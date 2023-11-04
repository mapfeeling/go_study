package _02311

import (
	"fmt"
	"testing"
)

func findMaximumXOR(nums []int) int {
	box := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cur := nums[i] ^ nums[j]
			if box < cur {
				box = cur
			}
		}
	}
	return box
}

func findMaximumXOROther(nums []int) (x int) {
	const highBit = 30 // 最高位的二进制位编号为 30
	for k := highBit; k >= 0; k-- {
		// 将所有的 pre^k(a_j) 放入哈希表中
		seen := map[int]bool{}
		for _, num := range nums {
			// 如果只想保留从最高位开始到第 k 个二进制位为止的部分
			// 只需将其右移 k 位
			seen[num>>k] = true
		}

		// 目前 x 包含从最高位开始到第 k+1 个二进制位为止的部分
		// 我们将 x 的第 k 个二进制位置为 1，即为 x = x*2+1
		xNext := x*2 + 1
		found := false

		// 枚举 i
		for _, num := range nums {
			if seen[num>>k^xNext] {
				found = true
				break
			}
		}

		if found {
			x = xNext
		} else {
			// 如果没有找到满足等式的 a_i 和 a_j，那么 x 的第 k 个二进制位只能为 0
			// 即为 x = x*2
			x = xNext - 1
		}
	}
	return
}

// 给你一个整数数组 nums ,返回 nums[i] XOR nums[j] 的最大运算结果,其中 0 ≤ i ≤ j < n
func TestFindMaximumXOR(t *testing.T) {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
	fmt.Println(findMaximumXOROther(nums))
}
