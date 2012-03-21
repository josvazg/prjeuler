package main

import (
	"fmt"
	"math"
    "time"
)

type long int64;

func maxFactor(target long) long {
	factor:=long(1);
    if((target%2)==0) {
    	factor=2;
    	for ;(target%2)==0; {
            //fmt.Println("*factor=",factor)
    		target=target/2;
    	}
    }
    i:=long(3);
    topper:=long(math.Sqrt(float64(target)));
    for ;target>1 && factor<=topper; {
    	if((target%i)==0) {
    		factor=i;
    		for ;(target%i)==0; {
    			target=target/factor;
                //fmt.Println("factor=",factor)
    		}
    		topper=long(math.Sqrt(float64(target)));
    	}
    	i+=2;
    }	
    if(target==1) {
    	return factor;
    }
    return target;
}

func main() {
    t:=time.Now()
	fmt.Println("maxFactor(500)=",maxFactor(500),"in",(time.Since(t)))
    t=time.Now()
    fmt.Println("maxFactor(600851475143)=",maxFactor(600851475143),"in",(time.Since(t)))
    t=time.Now()
    fmt.Println("maxFactor(600851475140)=",maxFactor(600851475140),"in",(time.Since(t)))
    t=time.Now()
    fmt.Println("maxFactor(60085147514)=",maxFactor(60085147514),"in",(time.Since(t)))
}


