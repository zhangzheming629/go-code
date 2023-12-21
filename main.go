package main

import (
	"fmt"
	"go-code/algo"
)

func main() {
	// IsPalindrome
	s := "A man, a plan, a canal: Panama"
	value := algo.IsPalindrome(s)
	fmt.Println(value)

	var arrs [5][6]int
	arrs[0][0] = 0
	arrs[0][1] = 0
	arrs[0][2] = 0
	arrs[0][3] = 0
	arrs[0][4] = 0
	arrs[0][5] = 0
	arrs[1][0] = 0
	arrs[1][1] = 0
	arrs[1][2] = 0
	arrs[1][3] = 0
	arrs[1][4] = 0
	arrs[1][5] = 0
	arrs[2][0] = 0
	arrs[2][1] = 0
	arrs[2][2] = 1
	arrs[2][3] = 0
	arrs[2][4] = 0
	arrs[2][5] = 0
	arrs[3][0] = 0
	arrs[3][1] = 0
	arrs[3][2] = 1
	arrs[3][3] = 1
	arrs[3][4] = 0
	arrs[3][5] = 0
	num := algo.Process(arrs)
	fmt.Println(num)
}
