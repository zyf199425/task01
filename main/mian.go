package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 只出现一次的数字
	nums := [5]int{4, 1, 2, 1, 2}
	res := singleNumber(nums[:])
	fmt.Println(res)

	// 回文数判断
	x := 121
	res2 := isPalindrome(x)
	fmt.Println(res2)

	// 有效括号
	s := "()[]{}"
	res3 := isValid(s)
	fmt.Println(res3)

	// 查找字符串数组中的最长公共前缀
	strs := []string{"flower", "flow", "flight"}
	res4 := longestCommonPrefix(strs)
	fmt.Println(res4)

	// 加一
	digits := []int{9, 9, 9, 9}
	res5 := plusOne(digits)
	fmt.Println(res5)

	// 两数之和
	nums1 := []int{2, 7, 11, 15}
	res7 := twoSum(nums1, 9)
	fmt.Println(res7)

}

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 回文数
// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	bytes := []byte(s)
	var left, right int = 0, len(s) - 1
	for left < right {
		if bytes[left] != bytes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	stack := Stack{}
	for index := range s {
		if isLeft(s[index]) {
			stack.push(s[index])
		} else {
			if stack.isEmpty() {
				return false
			}
			prev := stack.pop()
			if getLeft(s[index]) != prev {
				return false
			}
		}
	}
	return stack.isEmpty()
}
func isLeft(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}
func getLeft(s byte) byte {
	var res byte
	switch s {
	case ')':
		res = '('
	case '}':
		res = '{'
	case ']':
		res = '['
	}
	return res
}

type Stack []byte

func (s *Stack) push(v byte) {
	*s = append(*s, v)
}
func (s *Stack) pop() byte {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

// 查找字符串数组中的最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	end := 0
	b := strs[0][0]
outter:
	for ; end < len(strs[0]); end++ {
		for i := 1; i < len(strs); i++ {
			if end >= len(strs[i]) {
				break outter
			}
			if strs[i][end] != b {
				break outter
			}
		}
		if end+1 < len(strs[0]) {
			b = strs[0][end+1]
		}

	}
	return strs[0][0:end]
}

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			swap(nums, slow, fast)
		}
	}
	return slow + 1
}
func swap(nums []int, i int, j int) {
	var temp int = nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

// 加一
// https://leetcode.cn/problems/plus-one/description/
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		val := digits[i] + carry
		carry = val / 10
		digits[i] = val % 10
	}
	if carry > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// 两数之和
// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index := range nums {
		val, exits := m[target-nums[index]]
		if exits {
			return []int{val, index}
		}
		m[nums[index]] = index
	}
	return []int{-1, -1}
}
