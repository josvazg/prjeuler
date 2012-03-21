package main

import (
	"fmt"
	"math"
)

type long int64;

func maxFactor(target long) long {
	factor:=1;
    if((target%2)==0) {
    	factor=2;
    	for ;(target%2)==0; {
    		target=target/2;
    	}
    }
    i:=3;
    topper:=math.Sqrt(target);
    for ;target>1 && factor<=topper; {
    	if((target%i)==0) {
    		factor=i;
    		for ;(target%i)==0; {
    			target=target/factor;
    		}
    		topper=math.Sqrt(target);
    	}
    	i+=2;
    }	
    if(target==1) {
    	return factor;
    }
    return target;
}

func main() {
	fmt.Println("maxFactor(500)=",maxFactor())
}