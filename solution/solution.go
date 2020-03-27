package solution

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RemoveDuplicates 从排序数组中删除重复项
// 使用快慢机思想
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, 1

	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		}

		left++
		nums[left] = nums[right]
	}
	fmt.Println(nums[:left+1])
	return left + 1
}

// MaxProfit 买卖股票的最佳时机
// 将题目抽象化成数学问题，如果在Excel里画折线图就很容易理解，其实就是算上升期的数据总和。
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0 //总收益
	for i := 1; i < len(prices); i++ {
		dayPro := prices[i] - prices[i-1] //总是第一天买进第二天卖出，计算收益
		if dayPro > 0 {
			profit += dayPro //如果盈利就卖，并计算收益
		}
	}

	return profit
}

// RotateArr 旋转数组 冒泡
func RotateArr(nums []int, k int) {
	for j := 0; j < k; j++ {
		for i := len(nums) - 1; i > 0; i-- {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

// ContainsDuplicate 存在重复
// TODO 尝试不用sort包
func ContainsDuplicate(nums []int) bool {
	con := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := con[nums[i]]; ok {
			return true
		}

		con[nums[i]] = i
	}

	return false
}

// SingleNumber 只出现一次的数字
func SingleNumber(nums []int) int {
	start := 0
	for i := 0; i < len(nums); i++ {
		start = start ^ nums[i]
	}

	return start
}

// Intersect 两个数组的交集 II
func Intersect(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	flags := make([]bool, len2)

	var res []int

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if nums1[i] == nums2[j] && flags[j] == false {
				res = append(res, nums1[i])
				flags[j] = true
				break
			}
		}
	}

	return res

	// 有序的数组
	// sort.Ints(nums1)
	// sort.Ints(nums2)

	// len1 := len(nums1)
	// len2 := len(nums2)

	// var res []int
	// j := 0

	// for i := 0; i < len1; {
	// 	if i < len1 && j < len2 {
	// 		if nums1[i] == nums2[j] {
	// 			res = append(res, nums1[i])
	// 			i++
	// 			j++
	// 		} else if nums1[i] > nums2[j] {
	// 			j++
	// 		} else {
	// 			i++
	// 		}
	// 	} else {
	// 		break
	// 	}

	// }

	// return res

}

// PlusOne 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				h := []int{1}
				digits = append(h, digits...)
			}
			continue
		}
		break
	}

	return digits
}

// MoveZeroes 移动零
func MoveZeroes(nums []int) []int {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			flag++
		} else if flag != 0 {
			nums[i-flag] = nums[i]
			nums[i] = 0
		}
	}
	return nums
}

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	var res []int
	for i := 0; i < len(nums); i++ {
		flag := target - nums[i]
		if v, exsit := maps[flag]; exsit {
			res = append(res, v, i)
			break
		}
		maps[nums[i]] = i
	}
	return res
}

// IsValidSudoku 有效的数独
func IsValidSudoku(board [][]byte) bool {
	return false
}

// RotateMatrix 旋转图像
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	// 对角线翻转
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}

// ReverseString 反转字符串
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// ReverseInts 整数反转
func ReverseInts(x int) int {
	nums := 0
	for x != 0 {
		nums = nums*10 + x%10
		if nums >= 1<<31-1 || nums <= -1<<31 {
			return 0
		}

		x = x / 10
	}

	return nums
}

// FirstUniqChar 字符串中的第一个唯一字符
func FirstUniqChar(s string) int {
	maps := make([]int, 26)

	for i := 0; i < len(s); i++ {
		maps[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Println(s[i] - 'a')
		if maps[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1

}

// IsAnagram 有效的字母异位词
func IsAnagram(s string, t string) bool {
	if s == t {
		return true
	}

	if len(s) != len(t) {
		return false
	}

	nums := [26]rune{}

	for i, v := range s {
		nums[v-'a']++
		nums[rune(t[i])-'a']--
	}

	for _, v := range nums {
		if v != 0 {
			return false
		}
	}

	return true
}

// IsPalindrome 验证回文字符串
func IsPalindrome(s string) bool {
	len := len(s)

	if len == 0 {
		return true
	}
	s = strings.ToLower(s)
	left := 0
	right := len - 1
	for left < right {
		if !(s[left] >= 'a' && s[left] <= 'z' || s[left] >= 'A' && s[left] <= 'Z') {
			left++
		} else if !(s[right] > 'a' && s[right] < 'a' || s[right] > 'A' && s[right] < 'Z') {
			right--
		} else if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}

// MyAtoi 字符串转换整数 (atoi)
func MyAtoi(str string) int {
	if str == "" {
		return 0
	}
	for str[0] == ' ' {
		if len(str) > 1 {
			str = str[1:]
		} else {
			return 0
		}
	}
	flag := 1
	if str[0] == '-' {
		flag = -1
		if len(str) > 1 {
			str = str[1:]
		} else {
			return 0
		}

	}
	if str[0] == '+' && flag == 1 {
		if len(str) > 1 {
			str = str[1:]
		} else {
			return 0
		}
	}
	if len(str) == 0 {
		return 0
	}
	if str[0] < '0' || str[0] > '9' {
		return 0
	}
	index := 0
	for key, val := range str {
		if val < '0' || val > '9' {
			index = key
			break
		}
	}
	if index == 0 {
		index = len(str)
	}
	str = str[0:index]
	length := len(str)
	var res int64
	for i := 0; i < length; i++ {
		res *= 10
		if flag == 1 {
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			if -1*res < math.MinInt32 {
				return math.MinInt32
			}
		}
		res += int64(str[i] - '0')
		if flag == 1 {
			if res > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			if -1*res < math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	if flag == 1 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return int(res)
	}
	if -1*res < math.MinInt32 {
		return math.MinInt32
	}
	return int(-1 * res)

}

// StrStr 实现 strStr()
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteNode 删除链表中的节点
func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// RemoveNthFromEnd 删除链表的倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	slow := head

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	// 创建新链表
	var ret *ListNode
	for head != nil {
		// 临时保存下一个节点
		tmp := head.Next
		// 头结点重新指向新链表, 反转
		head.Next = ret
		// 新链表
		ret = head
		// 把下一个节点变成头节点,继续遍历
		head = tmp
	}

	return ret
}

// MergeTwoLists 合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建新链表
	var ret *ListNode
	return ret
}

// FizzBuzz FizzBuzz
func FizzBuzz(n int) []string {
	var str []string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			str = append(str, "FizzBuzz")
		} else if i%3 == 0 {
			str = append(str, "Fizz")
		} else if i%5 == 0 {
			str = append(str, "Buzz")
		} else {
			str = append(str, strconv.Itoa(i))
		}
	}

	return str
}
