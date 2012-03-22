package main

import (
	"fmt"
	"math"
	"time"
)

type long int64

func decomposeNAdd(primes []int, m long, a int) (long, []int) {
	for i := 1; i < len(primes); i++ { // decomponse a in primes (if possible)
		if a%primes[i] == 0 {
			a = a / primes[i]
			m = m * long(primes[i])
		}
	}
	// the remainder must be a prime, so we have to remember it
	if a > primes[len(primes)-1] {
		primes = append(primes, a)
		m = m * long(a) // update m with the newly found prime
	}
	return m, primes
}

func minDiv(n int) long {
	primes := []int{1}
	m := long(primes[0])
	for i := 2; i < n; i++ {
		if m%long(i) != 0 { // if m doesn't already divide by i...
			m, primes = decomposeNAdd(primes, m, i)
		}
	}
	return m
}

func isPrime(primes []int, a int) bool {
	for i := 1; i < len(primes); i++ {
		if a%primes[i] == 0 {
			return false
		}
	}
	return true
}

func nextPrime(primes []int) (int, []int) {
	for i := primes[len(primes)-1]; ; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
			//fmt.Println("primes=", primes)
			return i, primes
		}
	}
	return -1, primes
}

func minDivPE(n int) long {
	primes := []int{1, 2}
	m := long(1)
	check := true
	limit := int(math.Sqrt(float64(n)))
	for p := 2; p <= n; {
		p, primes = nextPrime(primes)
		a := 1
		if check {
			if p <= limit {
				a = int(math.Floor(math.Log(float64(n))) / math.Log(float64(p)))
			} else {
				check = false
			}
		}
		fmt.Println("p=", p, " check=", check)
		m = m * long(math.Pow(float64(p), float64(a)))
	}
	return m
}

func main() {
	fmt.Println("Project Eurler's Problem 5...")
	fmt.Println("Minimun number evenly divisible by all numbers of a sequence 1..N")

	t := time.Now()
	fmt.Println("1..10:", minDiv(10), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("PE 1..10:", minDivPE(10), "in", (time.Since(t)))
	t = time.Now()
	fmt.Println("1..20:", minDiv(20), "in", (time.Since(t)))
	/*t = time.Now()
	fmt.Println("PE 1..20:", minDivPE(20), "in", (time.Since(t)))*/

}

