package main

import (
	"fmt"
	"math"
	"time"
)

func isPalindrome(n int) bool {
	inv := 0
	for rest := n; rest > 1; rest = rest / 10 {
		inv = inv*10 + (rest % 10)
	}
	//fmt.Println(inv, "=", n, "?")
	return n == inv
}

func maxPalindrome1(digits int) int {
	top := int(math.Pow10(digits) - 1)
	min := int(math.Pow10(digits - 1))
	largestPalindrome := 0
	a := min
	//n := 0
	for ; a <= top; a++ {
		//n++
		b := min
		for ; b <= top; b++ {
			//n++
			c := a * b
			if c > largestPalindrome && isPalindrome(c) {
				largestPalindrome = c
			}
		}
	}
	//fmt.Println("n", n)
	return largestPalindrome
}

func maxPalindrome2(digits int) int {
	top := int(math.Pow10(digits) - 1)
	min := int(math.Pow10(digits - 1))
	largestPalindrome := 0
	//n := 0
	for a := top; a >= min; a-- {
		//n++
		b := top
		for ; b >= min; b-- {
			//n++
			c := a * b
			if c < largestPalindrome {
				break
			}
			if isPalindrome(c) {
				largestPalindrome = c
			}
		}
	}
	//fmt.Println("n", n)
	return largestPalindrome
}

func maxPalindrome3(digits int) int {
	top := int(math.Pow10(digits) - 1)
	min := int(math.Pow10(digits - 1))
	largestPalindrome := 0
	pow11 := top - top%11 //The largest number less than or equal top and divisible by 11
	//n := 0
	for a := top; a >= min; a-- {
		//n++
		b := pow11
		db := 11
		if a%11 == 0 {
			b = top
			db = 1
		}
		for ; b >= min; b -= db {
			//n++
			c := a * b
			if c < largestPalindrome {
				break
			}
			if isPalindrome(c) {
				largestPalindrome = c
			}
		}
	}
	//fmt.Println("n", n)
	return largestPalindrome
}

func main() {
	fmt.Println("Project Eurler's problem 4...")
	fmt.Println("Max Palindrome made up by multiplyng Ndigit numbers")

	t := time.Now()
	fmt.Println("1) 3digit numbers:", maxPalindrome1(3), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("2) 3digit numbers:", maxPalindrome2(3), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("3) 3digit numbers:", maxPalindrome3(3), "in", (time.Since(t)))

	t = time.Now()
	fmt.Println("1) 4digit numbers:", maxPalindrome1(4), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("2) 4digit numbers:", maxPalindrome2(4), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("3) 4digit numbers:", maxPalindrome3(4), "in", (time.Since(t)))

	t = time.Now()
	fmt.Println("1) 5digit numbers:", maxPalindrome1(5), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("2) 5digit numbers:", maxPalindrome2(5), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("3) 5digit numbers:", maxPalindrome3(5), "in", (time.Since(t)))
}

