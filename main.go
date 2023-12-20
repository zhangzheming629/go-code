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
	s = "race a car"
	value = algo.IsPalindrome(s)
	fmt.Println(value)
}
