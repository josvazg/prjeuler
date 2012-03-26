package main

import (
	"fmt"
	"math"
	"time"
)

type long int64

func isPrime(primes []long, a long) bool {
	for i := 1; i < len(primes); i++ {
		if a%primes[i] == 0 {
			return false
		}
	}
	return true
}

func primeAt(max int) long {
	primes:=[]long{1,2,3}
	for i := long(5); len(primes)<=max; i+=2 {
		if isPrime(primes, i) {
			primes = append(primes, i)
			//fmt.Println("prime=", i," len=",len(primes))
		}
	}
	return primes[len(primes)-1]
}

func isPrimePE(n int) bool {
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

func primeAtPE(n int) int {
	count:=1 //we  know that 2 is prime
	candidate:=3
	for count<n {
		if isPrimePE(candidate) { 
			count=count+1
			//fmt.Println(candidate,"is primer number",count)
		}
		if count<n {
			candidate=candidate+2
		}
	}
	return candidate
}

func main() {
	fmt.Println("Project Eurler's Problem 7...")
	fmt.Println("10001st prime")

	t := time.Now()
	fmt.Println("10001st prime:", primeAt(10001), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("10001st prime:", primeAtPE(10001), "in", (time.Since(t)))
}

