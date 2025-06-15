package main

import (
	"fmt"
	"slices"
	"strconv"
)

// 只出现一次的数字
func once(x []int) int {
	m := make(map[int]int)

	for _, v := range x {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 回文数
func isPalindrome(x int) bool {
	a := []byte(strconv.Itoa(x))
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	b, _ := strconv.Atoi(string(a))
	if x == b {
		return true
	}
	return false
}

// 有效的括号
func isValid(s string) bool {
	var a = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			a = append(a, '}')
		} else if s[i] == '(' {
			a = append(a, ')')
		} else if s[i] == '[' {
			a = append(a, ']')
		} else if len(a) != 0 && a[len(a)-1] == s[i] {
			a = a[:len(a)-1]
		} else {
			return false
		}
	}
	return len(a) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	s := []byte(strs[0])
	for _, str := range strs[1:] {
		var ti, li, minLen int
		b := []byte(str)
		if len(b) < len(s) {
			minLen = len(b)
		} else {
			minLen = len(s)
		}
		for i := 0; i < minLen; i++ {
			if s[i] != b[i] {
				li = i
				break
			} else {
				ti = i + 1
			}
		}
		if ti > li {
			s = s[:ti]
		} else {
			s = s[:li]
		}
	}
	return string(s)
}

// 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 加一
func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return digits
	}
	for i := l - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}

// 两数之和
func twoSum(nums []int, target int) []int {

	l := make([]int, 0)

	for i, num := range nums {
		for j, num2 := range nums {
			if num+num2 == target && i != j {
				return append(l, i, j)
			}
		}
	}
	return l
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 || len(intervals) == 0 {
		return intervals
	}
	slices.SortFunc(intervals, func(i, j []int) int {
		return i[0] - j[0]
	})
	start := intervals[0]
	ans := make([][]int, 0, len(intervals))
	for i, s := range intervals {
		if s[0] > start[1] {
			ans = append(ans, start)
			start = s
		} else {
			start[1] = max(start[1], s[1])
		}
		if i == len(intervals)-1 {
			ans = append(ans, start)
		}

	}
	return ans
}

func main() {
	// 回文数
	//fmt.Println(isPalindrome(121))
	//fmt.Println(isPalindrome(-121))
	//fmt.Println(isPalindrome(10))

	//只出现一次的数字
	//fmt.Println(once([]int{1, 2, 3, 4, 3, 2}))

	//有效的括号
	//fmt.Println(isValid("{}"))
	//fmt.Println(isValid("{]"))

	//fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	//fmt.Println(plusOne([]int{9}))

	//fmt.Println(twoSum([]int{3, 2, 4}, 6))

	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
