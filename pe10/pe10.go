package main

import (
	"fmt"
	"math"
	"time"
)

type long int64

func isPrime(n int) bool {
	if n==1 {
		return false
	} else if n<4 {
		return true  //2 and 3 are prime
	} else if n % 2==0 {
		return false
	} else if n<9 {
		return true    //we have already excluded 4,6 and 8.
	} else if n % 3==0 {
		return false
	} else {
		r:=int(math.Floor( math.Sqrt(float64(n)) ))  // n rounded to the greatest integer r so that r*r<=n
		f:=5
		for f<=r {
			if (n % f)==0 {
				return false
			}
			if (n %(f+2))==0 {
				return false
			}
			f=f+6
		}
		return true
	}
	return false
}

func primeSum(n int) long {
	count:=1 //we  know that 2 is prime
	candidate:=3
	sum:=long(2)
	for ;candidate<n;candidate=candidate+2 {
		if isPrime(candidate) { 
			count=count+1
			sum+=long(candidate)
			//fmt.Println(candidate,"is prime number",count)
		}
	}
	return sum
}

func main() {
	fmt.Println("Project Eurler's Problem 10...")
	fmt.Println("Find the sum of all the primes below 2000000")

	t := time.Now()
	fmt.Println("Sum of all the primes below 10:", primeSum(10), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("Sum of all the primes below 2000000:", primeSum(2000000), "in", (time.Since(t)))
}

