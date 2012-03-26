package main

import (
	"fmt"
	"math"
	"time"
)

func PithagoreanTripletProduct(sum int) int {
	for a:=1;a<sum;a++ {
		for b:=a;b<sum;b++ {
			fc:=math.Sqrt(float64((a*a)+(b*b)))
			c:=int(fc)
			if fc-float64(c)==0 {
				//fmt.Println("a=",a,"b=",b,"c=",c," sum=",(a+b+c))
				if a+b+c==sum {
					return a*b*c
				} else if a+b+c>sum {
					//fmt.Println("EXITS a=",a,"b=",b,"c=",c," sum=",(a+b+c))
					break;
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("Project Eurler's Problem 9...")
	fmt.Println("Pythagorean Triplet a^2+b^2=c^2 product a*b*c that a+b+c=1000")

	t := time.Now()
	fmt.Println("BPythagorean Triplet Product:", PithagoreanTripletProduct(1000), "in", (time.Since(t)))
}
