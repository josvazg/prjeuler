package main

import (
	"fmt"
	"time"
)

func minDiv(n int) int {
	primes:=[]int{1}
	m:=primes[0]
	for i:=2;i<n;i++ { 
		if m%i!=0 {
			a:=i
			//fmt.Println("a=",a)
			for j:=1;j<len(primes);j++ {
				if a%primes[j]==0 {
					a=a/primes[j]
					//fmt.Printf("*a=%v\n",a)
					for m%a!=0 && a%primes[j]==0 {
						a=a/primes[j]
						m=m*primes[j]
						fmt.Printf("x%v m=%v a=%v\n",primes[j],m,a)
					}
				}
			}
			if a>1 {
				primes=append(primes,a)
				m=m*a
				//fmt.Println("primes=",primes,"m=",m)
			}
		}
	}
	return m;
}

func main() {
	fmt.Println("Project Eurler's Problem 5...")
    fmt.Println("Minimun number evenly divisible by all numbers of a sequence 1..N")

    t:=time.Now()
    fmt.Println("1..10",minDiv(10),"in",(time.Since(t)))
    t=time.Now()
    fmt.Println("1..20",minDiv(20),"in",(time.Since(t)))
}