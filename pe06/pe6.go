package main

import (
	"fmt"
	"time"
)

type long int64

func difSqSum(max long) long {
	sum := (max * (max + 1)) / 2
	sqsum := sum * sum
	sumsq := long(0)
	for i := long(1); i <= max; i++ {
		sumsq += i * i
	}
	return sqsum - sumsq
}

func difSqSumO1(max long) long {
	sum := (max * (max + 1)) / 2
	sumsq := (2*max + 1) * (max + 1) * max / 6
	return (sum * sum) - sumsq
}

func main() {
	fmt.Println("Project Eurler's Problem 6...")
	fmt.Println("Difference between the sum of the squares and the square of the sums...")
	t := time.Now()
	fmt.Println("1-10:", difSqSum(10), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("O(1) 1-10:", difSqSumO1(10), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("1-100:", difSqSum(100), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("O(1) 1-100:", difSqSumO1(100), "in", (time.Since(t)))
}

