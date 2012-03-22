package main

import (
	"fmt"
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

func primes(max int) long {
	primes:=[]long{1,2,3}
	for i := long(5); len(primes)<=max; i+=2 {
		if isPrime(primes, i) {
			primes = append(primes, i)
			//fmt.Println("prime=", i," len=",len(primes))
		}
	}
	return primes[len(primes)-1]
}

func main() {
	fmt.Println("Project Eurler's Problem 7...")
	fmt.Println("10001st prime")

	t := time.Now()
	fmt.Println("10001st prime:", primes(10001), "in", (time.Since(t)))

}

