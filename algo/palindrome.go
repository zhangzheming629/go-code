package algo

import (
	"strings"
)

// 验证回文串: 给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。

// 本题中，将空字符串定义为有效的 回文串 。

// 示例 1:

// 输入: s = "A man, a plan, a canal: Panama"
// 输出: true
// 解释："amanaplanacanalpanama" 是回文串
// 示例 2:

// 输入: s = "race a car"
// 输出: false
// 解释："raceacar" 不是回文串

// 最简单的方法是对字符串 sss 进行一次遍历，并将其中的字母和数字字符进行保留，放在另一个字符串 sgood\textit{sgood}sgood 中。这样我们只需要判断 sgood\textit{sgood}sgood 是否是一个普通的回文串即可。

// 判断的方法有两种。

// 第一种是使用语言中的字符串翻转 API 得到 sgood\textit{sgood}sgood 的逆序字符串 sgood_rev\textit{sgood\_rev}sgood_rev，只要这两个字符串相同，那么 sgood\textit{sgood}sgood 就是回文串。

// 第二种是使用双指针。初始时，左右指针分别指向 sgood\textit{sgood}sgood 的两侧，随后我们不断地将这两个指针相向移动，每次移动一步，并判断这两个指针指向的字符是否相同。当这两个指针相遇时，就说明 sgood\textit{sgood}sgood 时回文串。

// 复杂度分析

// 时间复杂度：O(∣s∣)，其中 ∣s∣ 是字符串 sss 的长度。

// 空间复杂度：O(∣s∣)。由于我们需要将所有的字母和数字字符存放在另一个字符串中，在最坏情况下，新的字符串sgood 与原字符串 sss 完全相同，因此需要使用(∣s∣) 的空间。

func IsPalindrome(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
